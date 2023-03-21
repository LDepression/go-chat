/**
 * @Author: lenovo
 * @Description:
 * @File:  email
 * @Version: 1.0.0
 * @Date: 2023/03/20 22:55
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/logic"
	"go-chat/internal/model/request"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app"
	"go-chat/internal/pkg/app/errcode"
	"go.uber.org/zap"
)

type email struct {
}

// ExistEmail 判断邮箱是否已经被注册过了
// @Tags     email
// @Summary  判断邮箱是否被注册
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.CheckEmailExist  true  "email"
// @Success  200   {object}  common.State{}     "1001:参数有误 1003:系统错误 "
// @Router   /api/v1/email/check [post]
func (email) ExistEmail(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var reqEmail request.CheckEmailExist
	if err := ctx.ShouldBindJSON(&reqEmail); err != nil {
		zap.S().Infof("ctx.ShouldBindJSON(&reqEmail) failed,error=%v", err)
		return
	}
	ok, err := logic.Group.Email.CheckEmailIsUsed(ctx, reqEmail.Email)
	if err != nil {
		rly.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}
	reMap := map[string]interface{}{}
	if ok {
		reMap["status"] = "邮箱已经被注册过了"
	} else {
		reMap["status"] = "你可以使用该邮箱"
	}
	rly.Reply(nil, reMap)
	return
}

// SendEmail 发送邮件
// @Tags     email
// @Summary  发送邮件
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.SendEmail  true  "email"
// @Success  200   {object}  common.State{}     "1001:参数有误 1003:系统错误 3001:邮箱已经注册 3002:邮件发送频繁，请稍后再试"
// @Router   /api/v1/email/send [post]
func (email) SendEmail(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var reqEmail request.SendEmail
	if err := ctx.ShouldBindJSON(&reqEmail); err != nil {
		zap.S().Info("ctx.ShouldBindJSON(&reqEmail) failed,error=%v", err)
		return
	}
	if ok := logic.Group.Email.SendEmail(reqEmail.Email); ok != nil {
		rly.Reply(myerr.EmailSendTooMany)
		return
	}
	rly.Reply(nil)
	return
}
