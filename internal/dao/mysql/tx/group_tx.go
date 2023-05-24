/**
 * @Author: lenovo
 * @Description:
 * @File:  group_tx
 * @Version: 1.0.0
 * @Date: 2023/05/07 19:39
 */

package tx

import (
	"context"
	"go-chat/internal/dao"
	"go-chat/internal/global"
	"go-chat/internal/model/automigrate"
	"go-chat/internal/model/request"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type groupTX struct{}

func NewGroupTX() *groupTX {
	return &groupTX{}
}

// CreateGroupSetting 通过事务创建群组
func (groupTX) CreateGroupSetting(accountID uint, req request.CreateGroupReq) (uint, error) {
	tx := dao.Group.DB.Begin()
	relation := &automigrate.Relation{
		RelationType: "group",
		GroupType: automigrate.GroupType{
			Name:      req.Name,
			Signature: req.SigNature,
			Avatar:    global.Settings.Rule.DefaultAccountAvatar,
		},
	}
	if result := tx.Model(&automigrate.Relation{}).Create(relation); result.RowsAffected == 0 {
		tx.Rollback()
		return 0, gorm.ErrRecordNotFound
	}
	rID := relation.ID
	settingInfo := automigrate.Setting{
		AccountID:  accountID,
		RelationID: rID,
		IsSelf:     false,
		IsLeader:   true,
		IsShow:     true,
		IsPin:      false,
	}
	if result := tx.Model(&automigrate.Setting{}).CreateInBatches(&settingInfo, global.Settings.Rule.DefaultInsertDataNum); result.RowsAffected == 0 {
		tx.Rollback()

		return 0, result.Error
	}
	tx.Commit()
	//接下来将相关的东西保存到redis中去
	return relation.ID, dao.Group.Redis.AddRelationAccount(context.Background(), int64(rID), int64(accountID))
}

func (groupTX) Dissolve(relationID int64) error {
	tx := dao.Group.DB.Begin()
	if result := tx.Model(&automigrate.Setting{}).Where("relation_id = ?", relationID).Delete(&automigrate.Setting{}); result.RowsAffected == 0 {
		tx.Rollback()

		return result.Error
	}

	//还要去relation表中进行删除
	if result := tx.Model(&automigrate.Relation{}).Delete(&automigrate.Relation{}, relationID); result.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}
	//去redis中删除
	//去redis中找到对应的key
	accountIDs, err := dao.Group.Redis.GetAllAccountsByRelationID(context.Background(), relationID)
	if err != nil {
		tx.Rollback()
		zap.S().Infof("dao.Group.Redis.GetAllAccountsByRelationID failed,err:%v", err)
		return err
	}
	if err := dao.Group.Redis.DeleteRelationAccount(context.Background(), relationID, accountIDs...); err != nil {
		tx.Rollback()
		zap.S().Infof("dao.Group.Redis failed,err:%v", err)
		return err
	}
	tx.Commit()
	return nil
}

func (groupTX) AddAccounts2GroupWithTX(accountID int64, relationID int64, inviteIDs ...int64) error {
	//先将inviteID全部加入到数据库中去
	tx := dao.Group.DB.Begin()
	var examples []automigrate.Setting
	for _, inviteID := range inviteIDs {
		examples = append(examples, automigrate.Setting{
			AccountID:  uint(inviteID),
			RelationID: uint(relationID),
			IsSelf:     false,
			IsLeader:   false,
			IsShow:     true,
		})
	}
	if result := tx.Model(&automigrate.Setting{}).CreateInBatches(&examples, global.Settings.Rule.DefaultInsertDataNum); result.RowsAffected == 0 {
		tx.Rollback()

		return result.Error
	}
	//插入redis中去
	if err := dao.Group.Redis.AddRelationAccount(context.Background(), relationID, inviteIDs...); err != nil {
		zap.S().Infof(" dao.Group.Redis.AddRelationAccount failed: %v", err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

}

func (groupTX) GetGroupListTX(accountID int64) (*[]automigrate.Setting, error) {
	var ids []int64
	dao.Group.DB.Raw(`
		SELECT s.id 
		FROM
			settings s
			JOIN relations r ON 
			(
			s.relation_id IN ( 
				SELECT id FROM relations WHERE relation_type = "group" )
				AND s.relation_id = r.id
				)
		WHERE
			s.account_id = ? 
		ORDER BY
			s.last_show_time
`, accountID).Scan(&ids)
	if len(ids) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	//根据ids去进行相关的查询
	var GroupLists []automigrate.Setting
	if result := dao.Group.DB.Model(&automigrate.Setting{}).Preload("Relation").Find(&GroupLists, ids); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &GroupLists, nil
}

func (groupTX) TransferGroup(accountID int64, toID int64, relationID int64) error {
	tx := dao.Group.DB.Begin()
	if result := tx.Model(&automigrate.Setting{}).Where("account_id = ? and relation_id = ?", accountID, relationID).Update("is_leader", false); result.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}
	if result := tx.Model(&automigrate.Setting{}).Where("account_id = ? and relation_id = ?", toID, relationID).Update("is_leader", true); result.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}
	tx.Commit()
	return nil
}

func (groupTX) QuitGroup(accountID int64, relationID int64) error {
	//先去mysql中，将相关信息给删除
	tx := dao.Group.DB.Begin()
	if result := tx.Model(&automigrate.Setting{}).Where("account_id = ? AND relation_id =?", accountID, relationID).Delete(&automigrate.Setting{}); result.RowsAffected == 0 {
		tx.Rollback()
		zap.S().Infof("delete failed ")
		return gorm.ErrRecordNotFound
	}
	//去redis中将相关的信息进行删除
	if err := dao.Group.Redis.DeleteRelationAccount(context.Background(), relationID, accountID); err != nil {
		zap.S().Infof("dao.Group.Redis.DeleteRelationAccount failed err:%v", err)
		return err
	}
	tx.Commit()
	return nil
}
