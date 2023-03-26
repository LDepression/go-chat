/**
 * @Author: lyc
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/03/20 18:46
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/api/base"
	"go-chat/internal/logic"
	"go-chat/internal/model/request"
	"go-chat/internal/pkg/app"
	"go-chat/internal/pkg/app/errcode"
	"go.uber.org/zap"
)

type user struct {
}

func NewUser() *user {
	return &user{}
}

// Register 用户注册
// @Tags     register
// @Summary  用户注册
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.Register                   true  "用户注册信息"
// @Success  200   {object}  common.State{data=reply.Register}  "1001:参数有误 1003:系统错误 3001:邮箱已经注册 "
// @Router   /api/v1/user/register [post]
func (user) Register(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var reqRegister request.Register
	if err := ctx.ShouldBindJSON(&reqRegister); err != nil {
		zap.S().Info("ctx.ShouldBindJSON(&reqRegister) failed: %v", err)
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	replyInfo, err := logic.Group.User.Register(ctx, reqRegister.Mobile, reqRegister.Email, reqRegister.Password, reqRegister.EmailCode)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, replyInfo)
}

// Login 用户登录
// @Tags     login
// @Summary  用户登录
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.Login                   true  "用户登录信息"
// @Success  200   {object}  common.State{data=reply.LoginReply}  "1001:参数有误 1003:系统错误 3001:邮箱已经注册 "
// @Router   /api/v1/user/login [post]
func (user) Login(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var reqLogin request.Login
	if err := ctx.ShouldBindJSON(&reqLogin); err != nil {
		zap.S().Info("ctx.ShouldBindJSON(&reqLogin) failed: %v", err)
		base.HandleValidatorError(ctx, err)
		return
	}
	replyInfo, err := logic.Group.User.Login(ctx, reqLogin)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, replyInfo)
}

// ModifyPassword 用户更改密码
// @Tags     ModifyPassword
// @Summary  用户更改密码
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.ReqModifyPassword                   true  "用户登录信息"
// @Success  200   {object}  common.State{}  "1001:参数有误 1003:系统错误 3001:邮箱已经注册 "
// @Router   /api/v1/user/modifyPassword [post]
func (user) ModifyPassword(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var reqModify request.ReqModifyPassword
	if err := ctx.ShouldBindJSON(&reqModify); err != nil {
		zap.S().Info("ctx.ShouldBindJSON(&reqModify) failed: %v", err)
		base.HandleValidatorError(ctx, err)
		return
	}
	if err := logic.Group.User.ModifyPassword(ctx, reqModify); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

// Logout 用户登出
// @Tags     Logout
// @Summary  用户登出
// @accept   application/json
// @Produce  application/json
// @Param    data  body                        true  "用户登录信息"
// @Success  200   {object}  common.State{}  "1001:参数有误 1003:系统错误 3001:邮箱已经注册 "
// @Router   /api/v1/user/logout [post]
func (user) Logout(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	if err := logic.Group.User.Logout(ctx); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, gin.H{
		"msg": "登出成功",
	})
}
