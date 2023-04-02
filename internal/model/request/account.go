package request

import (
	"go-chat/internal/model/common"
)

type AccountRequestInfo struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	Signature string `json:"signature"`
	Avatar    string `json:"avatar"`
	Gender    int    `json:"gender"`
}

type GetAccountByID struct {
	AccountID uint `json:"account_id" form:"account_id" binding:"required,gte=1"`
}

type GetAccountsByName struct {
	AccountName string `json:"account_name" form:"account_name" binding:"required"`
	common.Pager
}

type UpdateAccount struct {
	AccountID uint   `json:"account_id" form:"account_id" binding:"required,gte=1"`
	Name      string `json:"name,omitempty" form:"name"`
	Signature string `json:"signature,omitempty" form:"signature"`
	Avatar    string `json:"avatar,omitempty" form:"avatar"`
	Gender    int    `json:"gender,omitempty" form:"gender"`
}
