/**
 * @Author: lenovo
 * @Description:
 * @File:  account
 * @Version: 1.0.0
 * @Date: 2023/03/28 16:36
 */

package request

type GenderType int

var (
	MALE   GenderType = 1
	FEMALE GenderType = 2
)

type CreateAccountReq struct {
	ID        int64       `json:"ID"`
	Name      string      `json:"name" binding:"required"`         //昵称
	Avatar    string      `json:"avatar" binding:"required"`       //头像图片
	Signature string      `json:"sigNature"`                       //个性签名
	Gender    *GenderType `json:"gender" binding:"required,min=0"` //性别
}

type DeleteAccountReq struct {
	AccountID int64 `json:"accountID"` //删除账号的ID
}
