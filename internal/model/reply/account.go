/**
 * @Author: lenovo
 * @Description:
 * @File:  account
 * @Version: 1.0.0
 * @Date: 2023/03/28 22:49
 */

package reply

import (
	"go-chat/internal/model/common"
	"time"
)

type CreateAccountReply struct {
	AccountID int64  `json:"accountID"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`

	Token common.Token `json:"token"`
}

type AccountInfoReply struct {
	AccountID int64  `json:"accountID"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Avatar    string `json:"avatar"`
	Signature string `json:"signature"`
}
type TotalAccountsReply struct {
	Total        int                `json:"total"`
	AccountInfos []AccountInfoReply `json:"accountInfos"`
}

type AccountInfo struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	Signature string    `json:"signature"`
	Avatar    string    `json:"avatar"`
	Gender    string    `json:"gender"`
}

type GetAccountByID struct {
	AccountInfo
}

type GetAccountsByName struct {
	AccountInfos []*AccountInfo
	Total        int64
}

type GetAccountsByUserID struct {
	AccountInfos []*AccountInfo
	Total        int64
}
