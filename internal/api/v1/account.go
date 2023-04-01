/**
 * @Author: lenovo
 * @Description:
 * @File:  account
 * @Version: 1.0.0
 * @Date: 2023/03/28 16:44
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/api/base"
	"go-chat/internal/logic"
	"go-chat/internal/middleware"
	"go-chat/internal/model"
	"go-chat/internal/model/request"
	"go-chat/internal/pkg/app"
	"go-chat/internal/pkg/app/errcode"
	"go-chat/internal/pkg/utils"
	"go.uber.org/zap"
)

type account struct{}

// CreateAccount 创建账户
// @Tags     account
// @Summary  创建账户
// @accept   application/json
// @Produce  application/json
// @Param 	Authorization 	header 	string 	true 	"x-token 用户令牌"
// @Param   data  body      request.CreateAccountReq  true  "email"
// @Success  200   {object}  common.State{data=reply.CreateAccountReply}     "1001:参数有误 1003:系统错误 "
// @Router   /api/v1/account/createAccount [post]
func (account) CreateAccount(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var req request.CreateAccountReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.HandleValidatorError(ctx, err)
		zap.S().Infof("ctx.ShouldBindJSON(&req) failed: %v", err)
		return
	}
	replyInfo, err := logic.Group.Account.CreateAccount(ctx, req)
	if err != nil {
		rly.Reply(err)
		zap.S().Info("logic.Group.Account.CreateAccount failed", zap.Any("err", err))
		return
	}
	rly.Reply(nil, replyInfo)
}

// DeleteAccount 删除账户
// @Tags     account
// @Summary  删除账户
// @accept   application/json
// @Produce  application/json
// @Param 	Authorization 	header 	string 	true 	"x-token 用户令牌"
// @Param   data  query      request.DeleteAccountReq  true  "account_id"
// @Success  200   {object}  common.State{}     "1001:参数有误 1003:系统错误 "
// @Router   /api/v1/account/deleteAccount/{id} [delete]
func (account) DeleteAccount(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	accountID := ctx.Param("id")
	id := utils.StringToIDMust(accountID)
	if err := logic.Group.Account.DeleteAccount(ctx, id); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

// GetToken 获取账户令牌
// @Tags     account
// @Summary  获取账户令牌
// @accept   application/json
// @Produce  application/json
// @Param 	Authorization 	header 	string 	true 	"x-token 用户令牌"
// @Param   data  query      request.DeleteAccountReq  true  "account_id"
// @Success  200   {object}  common.State{data=common.Token}     "1001:参数有误 1003:系统错误 "
// @Router   /api/v1/account/getToken/{id} [get]
func (account) GetToken(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	accountIDStr := ctx.Param("id")
	id := utils.StringToIDMust(accountIDStr)
	replyInfo, err := logic.Group.Account.GetToken(ctx, id)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, replyInfo)
}

// GetAccountsByUserID 获取用户的所有账号
// @Tags     account
// @Summary  获取用户的所有账号
// @accept   application/json
// @Produce  application/json
// @Param 	Authorization 	header 	string 	true 	"x-token 用户令牌"
// @Success  200   {object}  common.State{data=reply.TotalAccountsReply}     "1001:参数有误 1003:系统错误 "
// @Router   /api/v1/account//infos/user [get]
func (account) GetAccountsByUserID(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	content, ok := middleware.GetContent(ctx)
	if !ok || content.Type != model.UserToken {
		rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
		return
	}
	AccountInfosReply, err := logic.Group.Account.GetAccountsByUserID(content.ID)
	rly.Reply(err, AccountInfosReply)
}
