/**
 * @Author: lenovo
 * @Description:
 * @File:  account
 * @Version: 1.0.0
 * @Date: 2023/03/28 16:36
 */

package request

type GenderType int

type CreateAccountReq struct {
	Name      string `json:"name" binding:"required"`   //昵称
	Avatar    string `json:"avatar" binding:"required"` //头像图片
	Signature string `json:"sigNature"`                 //个性签名
	Gender    int    `json:"gender" `
}
