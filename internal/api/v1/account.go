package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-chat/internal/api/base"
	"go-chat/internal/global"
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

func NewAccount() *account {
	return &account{}
}

// GetAccountByID
// @Tags     account
// @Summary  获取账户信息
// @accept   application/form-data
// @Produce  application/json
// @Param    Authorization  header    string               true "x-token 用户令牌"
// @Param    data           query     request.GetAccountByID                   true  "账号信息"
// @Success  200            {object}  common.State{data=reply.GetAccountByID}  "1001:参数有误 1003:系统错误 2009:权限不足 2007:身份不存在 2008:身份验证失败 2010:账号不存在"
// @Router   /api/account/info [get]
func (account) GetAccountByID(c *gin.Context) {
	res := app.NewResponse(c)
	params := &request.GetAccountByID{}
	if err := c.ShouldBindQuery(params); err != nil {
		zap.S().Errorf("&request.GetAccountByID{} c.ShouldBindJSON(params) failed: %v", err)
		res.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	zap.S().Infof("params.AccountID:%v", params.AccountID)
	fmt.Println("params.AccountID:", params.AccountID)
	result, err := logic.Group.Account.GetAccountByID(c, params.AccountID)
	if err != nil {
		res.Reply(err)
		return
	}
	res.Reply(nil, result)
}

// GetAccountsByName
// @Tags     account
// @Summary  通过昵称模糊查找账户
// @accept   application/form-data
// @Produce  application/json
// @Param    Authorization  header    string            true "x-token 用户令牌"
// @Param    data           query     request.GetAccountsByName                   true  "账号信息"
// @Success  200            {object}  common.State{data=reply.GetAccountsByName}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在"
// @Router   /api/account/infos/name [get]
func (account) GetAccountsByName(c *gin.Context) {
	res := app.NewResponse(c)
	params := &request.GetAccountsByName{}
	if err := c.ShouldBindJSON(params); err != nil {
		zap.S().Errorf("&request.GetAccountByName{} c.ShouldBindJSON(params) failed: %v", err)
		res.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	limit, offset := global.Pager.GetPageSizeAndOffset(c)
	result, err := logic.Group.Account.GetAccountsByName(params.AccountName, limit, offset)

	res.ReplyList(err, result.Total, result.AccountInfos)
}

// CreateAccount 创建账户
// @Tags     account
// @Summary  创建账户
// @Param    Authorization  header    string          true "x-token 用户令牌"
// @Success  200            {object}  common.State{data=reply.GetAccountsByUserID}  "1003:系统错误 2008:身份验证失败 2010:账号不存在"
// @Router   /api/account/infos/user [get]
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
		zap.S().Errorf("logic.Group.Account.CreateAccount failed, err:%v", err)
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
	AccountInfosReply, err := logic.Group.Account.GetAccountsByUserID(int64(content.ID))
	rly.Reply(err, AccountInfosReply)
}

// UpdateAccount
// @Tags     account
// @Summary  更新账户信息
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string        false "x-token 用户令牌"
// @Param    data           body      request.UpdateAccount  true  "账号信息"
// @Success  200            {object}  common.State{}         "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败"
// @Router   /api/account/update [put]
func (account) UpdateAccount(c *gin.Context) {
	res := app.NewResponse(c)
	params := &request.UpdateAccount{}
	if err := c.ShouldBindJSON(params); err != nil {
		zap.S().Errorf("&request.UpdateAccount{} c.ShouldBindJSON(params) failed: %v", err)
		res.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := middleware.GetContent(c)
	if !ok || content.Type != model.AccountToken {
		res.Reply(errcode.AuthNotExist)
		return
	}
	err := logic.Group.Account.UpdateAccount(c, content.ID, params.Name, params.Signature, params.Avatar, string(params.Gender))
	res.Reply(err)
}
