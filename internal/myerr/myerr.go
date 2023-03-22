/**
 * @Author: lenovo
 * @Description:
 * @File:  myerr
 * @Version: 1.0.0
 * @Date: 2023/03/20 21:29
 */

package myerr

import "go-chat/internal/pkg/app/errcode"

var (
	EmailExists      = errcode.NewErr(3001, "邮箱已经注册")
	EmailSendTooMany = errcode.NewErr(3002, "发送次数过多")
	EmailCodeInvalid = errcode.NewErr(3003, "验证码失效或错误")
)
