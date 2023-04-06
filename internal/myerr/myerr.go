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
	EmailNotFound    = errcode.NewErr(3004, "邮箱不存在")
	PasswordError    = errcode.NewErr(3005, "密码错误")
	PasswordInvalid  = errcode.NewErr(3006, "密码不能为空")
	ChoiceNotFound   = errcode.NewErr(3008, "登录选项错误")
	UserNotFound     = errcode.NewErr(3009, "请先注册或登录")
	TokenInValid     = errcode.NewErr(3010, "token失效了")
	TokenNotFound    = errcode.NewErr(3011, "token不存在")
	UserNotExist     = errcode.NewErr(3012, "用户不存在")
	UserExist        = errcode.NewErr(3013, "用户已经登录了")
	AuthFailed       = errcode.NewErr(3014, "身份认证失败")

	AccountNotExist     = errcode.NewErr(4001, "账户不存在")
	ApplicationNotFound = errcode.NewErr(5001, "好友申请不存在")
)
