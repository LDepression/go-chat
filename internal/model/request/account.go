package request

import "go-chat/internal/model/common"

type GetAccountByID struct {
	AccountID int64 `json:"account_id" form:"account_id" binding:"required,gte=1"`
}

type GetAccountListByUserID struct {
}

type GetAccountsByName struct {
	AccountName string `json:"account_name" form:"account_name" binding:"required"`
	common.Pager
}
