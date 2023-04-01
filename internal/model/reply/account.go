/**
 * @Author: lenovo
 * @Description:
 * @File:  account
 * @Version: 1.0.0
 * @Date: 2023/03/28 22:49
 */

package reply

import "go-chat/internal/model/common"

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
