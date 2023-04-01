/**
 * @Author: lenovo
 * @Description:
 * @File:  account
 * @Version: 1.0.0
 * @Date: 2023/03/28 22:37
 */

package query

import (
	"go-chat/internal/dao"
	"go-chat/internal/model/automigrate"
)

type account struct{}

func NewAccount() *account {
	return &account{}
}

func (account) CheckAccountInfosByUserID(userID int64) ([]automigrate.Account, error) {
	var accounts []automigrate.Account
	if result := dao.Group.DB.Where("user_id = ?", userID).Find(&accounts); result.RowsAffected == 0 {
		return accounts, result.Error
	}
	return accounts, nil
}

func (account) CheckAccountInfoByAccountID(accountID int64) (automigrate.Account, error) {
	var accountInfo automigrate.Account
	if result := dao.Group.DB.Where("ID = ?", accountID).Find(&accountInfo); result.RowsAffected == 0 {
		return accountInfo, result.Error
	}
	return accountInfo, nil
}

func (account) GetAccountsByUserID(userID int64) ([]automigrate.Account, error) {
	var accountInfos []automigrate.Account
	result := dao.Group.DB.Where("user_id = ?", userID).Find(&accountInfos)
	return accountInfos, result.Error
}
