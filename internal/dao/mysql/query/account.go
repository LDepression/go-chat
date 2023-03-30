package query

import (
	"go-chat/internal/dao"
	"go-chat/internal/model/automigrate"
)

type qAccount struct {
}

func NewQueryAccount() *qAccount {
	return &qAccount{}
}

func (qAccount) GetAccountByID(AccountID int64) (*automigrate.Account, error) {
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
