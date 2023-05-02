/**
 * @Author: lenovo
 * @Description:
 * @File:  account
 * @Version: 1.0.0
 * @Date: 2023/03/28 22:37
 */

package query

import (
	"fmt"
	"go-chat/internal/dao"
	"go-chat/internal/model/automigrate"
)

type qAccount struct {
}

func NewQueryAccount() *qAccount {
	return &qAccount{}
}

func (qAccount) GetAccountByID(AccountID uint) (*automigrate.Account, error) {
	var accountInfo automigrate.Account

	if result := dao.Group.DB.Model(&automigrate.Account{}).Where("id = ?", AccountID).First(&accountInfo); result.Error != nil {
		return nil, result.Error
	}
	return &accountInfo, nil
}

func (qAccount) GetAccountsByName(AccountName string, limit, offset int32) ([]*automigrate.Account, int64, error) {
	var accountInfos []*automigrate.Account
	var totalCount int64
	if result := dao.Group.DB.Model(&automigrate.Account{}).
		Where("name like ?", "%"+AccountName+"%").
		Offset(int(offset)).Limit(int(limit)).
		Find(&accountInfos); result.Error != nil {
		return nil, 0, result.Error
	} else {
		totalCount = result.RowsAffected
	}

	return accountInfos, totalCount, nil
}

func (qAccount) UpdateAccount(accountID uint, name, signature, avatar, gender string) error {
	updateFields := map[string]interface{}{
		"name":      name,
		"signature": signature,
		"avatar":    avatar,
		"gender":    gender,
	}
	for k := range updateFields {
		if updateFields[k] == "" {
			delete(updateFields, k)
		}
	}
	if result := dao.Group.DB.Model(&automigrate.Account{}).Where("id = ?", accountID).Updates(updateFields); result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("no account with ID %d found", accountID)
	}
	return nil
}

func (qAccount) CheckAccountInfosByUserID(userID int64) ([]automigrate.Account, error) {
	var accounts []automigrate.Account
	if result := dao.Group.DB.Where("user_id = ?", userID).Find(&accounts); result.RowsAffected == 0 {
		return accounts, result.Error
	}
	return accounts, nil
}

func (qAccount) CheckAccountInfoByAccountID(accountID int64) (automigrate.Account, error) {
	var accountInfo automigrate.Account
	if result := dao.Group.DB.Where("ID = ?", accountID).Find(&accountInfo); result.RowsAffected == 0 {
		return accountInfo, result.Error
	}
	return accountInfo, nil
}

func (qAccount) GetAccountsByUserID(userID int64) ([]automigrate.Account, error) {
	var accountInfos []automigrate.Account
	result := dao.Group.DB.Where("user_id = ?", userID).Find(&accountInfos)
	return accountInfos, result.Error
}
