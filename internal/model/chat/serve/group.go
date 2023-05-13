/**
 * @Author: lenovo
 * @Description:
 * @File:  group
 * @Version: 1.0.0
 * @Date: 2023/05/08 22:25
 */

package serve

type InviteNewPerson struct {
	Encoding  string `json:"Encoding"` //进行加密之后的token
	AccountID int64  `json:"AccountID"`
}

type DissolveGroup struct {
	Encoding  string `json:"Encoding"`
	AccountID int64  `json:"AccountID"`
}

type TransferGroup struct {
	Encoding  string `json:"Encoding"`
	AccountID int64  `json:"AccountID"`
}

type QuitGroup struct {
	Encoding  string `json:"Encoding"`
	AccountID int64  `json:"AccountID"`
}
