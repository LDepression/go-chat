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
// @Tags     user
// @Summary  用户注册
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.Register                   true  "用户注册信息"
// @Success  200   {object}  common.State{data=reply.Register}  "1001:参数有误 1003:系统错误 2004:邮箱验证码校验失败 2006:邮箱已经注册 "
// @Router   /api/v1/user/register [post]
func (user) Register(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var reqRegister request.Register
	if err := ctx.ShouldBindJSON(&reqRegister); err != nil {
		zap.S().Info("ctx.ShouldBindJSON(&reqRegister) failed: %v", err)
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := logic.Group.User.Register(ctx, reqRegister.Mobile, reqRegister.Email, reqRegister.Password, reqRegister.EmailCode); err != nil {
		rly.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}
	rly.Reply(nil)
}
