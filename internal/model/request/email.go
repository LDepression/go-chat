/**
 * @Author: lenovo
 * @Description:
 * @File:  emai;
 * @Version: 1.0.0
 * @Date: 2023/03/20 22:43
 */

package request

type SendEmail struct {
	Email string `json:"email" binding:"required,email"` //发送邮箱的验证码
}

type CheckEmailExist struct {
	Email string `json:"email" binding:"required,email"` //检查是否存在的邮箱
}
