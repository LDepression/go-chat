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
	if result := dao.Group.DB.Model(&automigrate.Account{}).Where("name like ?", "%"+AccountName+"%").Offset(int(offset)).Limit(int(limit)).Find(&accountInfos); result.Error != nil {
		return nil, 0, result.Error
	} else {
		totalCount = result.RowsAffected
	}

	return accountInfos, totalCount, nil
}

func (qAccount) GetAccountsByUserID(userID uint) ([]*automigrate.Account, int64, error) {
	var user automigrate.User
	if result := dao.Group.DB.Preload("Accounts").First(&user, userID); result.Error != nil {
		return nil, 0, result.Error
	}
	return user.Accounts, int64(len(user.Accounts)), nil
}

func (qAccount) GetUserByAccountID(accountID int64) (*automigrate.User, error) {
	var account automigrate.Account
	if result := dao.Group.DB.Preload("User").First(&account, accountID); result.Error != nil {
		return nil, result.Error
	}
	return account.User, nil
}

func (qAccount) UpdateAccount(accountID uint, name, signature, avatar string, gender int) error {
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
