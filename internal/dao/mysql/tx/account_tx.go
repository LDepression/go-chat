/**
 * @Author: lenovo
 * @Description:
 * @File:  account_tx
 * @Version: 1.0.0
 * @Date: 2023/03/28 23:04
 */

package tx

import (
	"context"
	"errors"
	"go-chat/internal/dao"
	"go-chat/internal/global"
	"go-chat/internal/model/automigrate"
	"go-chat/internal/model/request"
)

var (
	AccountMaxNums = global.Settings.Rule.AccountMaxNums
)

var (
	ErrExceedAccountNums       = errors.New("超过最大创建账户的数量了")
	ErrUserNameHasBeenRegister = errors.New("用户名已经被该用户注册过了")
	ErrAccountIsLeader         = errors.New("该账户是群主")
	ErrAccountNotExist         = errors.New("用户不存在")
)

type AccountTx struct {
}

func NewAccountTX() *AccountTx {
	return &AccountTx{}
}
func (AccountTx) CreateAccountWithTX(ctx context.Context, userID int64, req request.CreateAccountReq) error {
	var accounts []automigrate.Account
	var user automigrate.User
	tx := dao.Group.DB.Begin()
	if result := tx.Model(&automigrate.User{}).Find(&user, userID); result.RowsAffected == 0 {
		tx.Rollback()
		return result.Error
	}
	if len(accounts) > AccountMaxNums {
		tx.Rollback()

		return ErrExceedAccountNums
	}
	//在去判断一下用户名是否已经被用过
	if result := tx.Where(&automigrate.Account{UserID: uint(userID), Name: req.Name}).Find(&accounts); result.RowsAffected != 0 {
		tx.Rollback()
		return ErrUserNameHasBeenRegister
	}
	var AccountInfo automigrate.Account
	AccountInfo.ID = int32(req.ID)
	AccountInfo.Name = req.Name
	AccountInfo.UserID = uint(userID)
	AccountInfo.Avatar = req.Avatar
	AccountInfo.Gender = string(*req.Gender)
	AccountInfo.Signature = req.Signature
	if result := tx.Create(&AccountInfo); result.RowsAffected == 0 {
		tx.Rollback()

		return result.Error
	}
	var relation automigrate.Relation
	//接下来去创建关系,因为是创建账户,所以存一个单向关系就好了
	relation.RelationType = "friend"
	relation.FriendType = automigrate.FriendType{
		AccountID1: int64(AccountInfo.ID),
		AccountID2: int64(AccountInfo.ID),
	}
	if result := tx.Create(&relation); result.RowsAffected == 0 {
		tx.Rollback()

		return result.Error
	}
	var setting automigrate.Setting
	setting.RelationID = relation.ID
	setting.AccountID = uint(AccountInfo.ID)
	setting.IsSelf = true
	if result := tx.Create(&setting); result.RowsAffected == 0 {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	if err := dao.Group.Redis.AddRelationAccount(ctx, int64(relation.ID), []int64{int64(AccountInfo.ID)}); err != nil {
		return err
	}
	return nil
}

func (AccountTx) DeleteAccountWithTX(ctx context.Context, accountID int64) error {
	tx := dao.Group.DB.Begin()
	//先去判断一下 ,该用户是否是群主,如果是的话,不能删除
	if result := tx.Model(&automigrate.Setting{}).Where(automigrate.Setting{
		AccountID: uint(accountID),
		IsLeader:  true,
	}).Find(&automigrate.Setting{}); result.RowsAffected != 0 {
		tx.Rollback()
		return ErrAccountIsLeader
	}
	//看一下accountID是否存在于account表中
	if result := tx.Model(&automigrate.Account{}).Find(&automigrate.Account{}, accountID); result.RowsAffected == 0 {
		tx.Rollback()
		return ErrAccountNotExist
	}
	//再在Account表里面通过AccountID将这条数据删除

	var accounts []automigrate.Account

	var accountIDs []int64
	if result := tx.Model(&automigrate.Account{}).Where(&automigrate.Account{
		BaseModel: automigrate.BaseModel{ID: int32(accountID)},
	}).Delete(&accounts); result.RowsAffected == 0 {
		tx.Rollback()
		return result.Error
	}
	for _, account := range accounts {
		accountIDs = append(accountIDs, int64(account.ID))
	}

	//再将relation表删除中通过accountID字段进行删除
	var relation automigrate.Relation
	if result := tx.Model(&automigrate.Relation{}).Where("JSON_EXTRACT(friend_type, '$.AccountID1') = ?", accountID).Or(
		"JSON_EXTRACT(friend_type, '$.AccountID2') = ?", accountID,
	).Delete(&relation); result.RowsAffected == 0 {
		tx.Rollback()
		return result.Error
	}

	//最后在accountID所在的群,将account字段删除
	var settings []automigrate.Setting
	var oldSettings []automigrate.Setting
	if result := tx.Model(&automigrate.Setting{}).Where("account_id = ?", accountID).Find(&oldSettings); result.RowsAffected == 0 {
		return result.Error
	}
	if result := tx.Model(&automigrate.Setting{}).Where("account_id =?", accountID).Delete(&settings); result.RowsAffected == 0 {
		tx.Rollback()
		return result.Error
	}
	var relationIDs []int64
	for _, setting := range oldSettings {
		relationIDs = append(relationIDs, int64(setting.RelationID))
	}
	tx.Commit()

	//在一个群删除多个人
	if err := dao.Group.Redis.DeleteRelationAccount(ctx, int64(relation.ID), accountIDs...); err != nil {
		tx.Rollback()
		return err
	}

	//在多个群将一个人删除
	if err := dao.Group.Redis.DeleteAccountByRelations(ctx, accountID, relationIDs); err != nil {
		return err
	}
	return nil
}
