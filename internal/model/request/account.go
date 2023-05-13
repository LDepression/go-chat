package request

import (
	"go-chat/internal/model/common"
)

type GenderType string

const (
	MALE   GenderType = "male"
	FEMALE GenderType = "female"
)

type CreateAccountReq struct {
	ID        int64      `json:"ID"`
	Name      string     `json:"name" binding:"required"`             //昵称
	Signature string     `json:"sigNature"`                           //个性签名
	Gender    GenderType `json:"gender" binding:"required,oneof=男 女"` //性别
}

type DeleteAccountReq struct {
	AccountID int64 `json:"accountID"` //删除账号的ID
}
type GetAccountByID struct {
	AccountID uint `json:"account_id" form:"account_id" `
}

type GetAccountsByName struct {
	AccountName string `json:"account_name" form:"account_name" `
	common.Pager
}

type UpdateAccount struct {
	Name      string     `json:"name,omitempty" form:"name"`
	Signature string     `json:"signature,omitempty" form:"signature"`
	Avatar    string     `json:"avatar,omitempty" form:"avatar"`
	Gender    GenderType `json:"gender,omitempty" form:"gender"`
}
