package v1

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/global"
	"go-chat/internal/logic"
	"go-chat/internal/middleware"
	"go-chat/internal/model"
	"go-chat/internal/model/request"
	"go-chat/internal/pkg/app"
	"go-chat/internal/pkg/app/errcode"
	"go.uber.org/zap"
)

type account struct {
}

func NewAccount() *account {
	return &account{}
}

// GetAccountByID
// @Tags     account
// @Summary  获取账户信息
// @accept   application/form-data
// @Produce  application/json
// @Param    Authorization  header    string               false "x-token 用户令牌"
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
	//content, ok := middleware.GetPayLoad(c)
	//if !ok || content.Type != model.AccountToken {
	//	res.Reply(errcode.AuthNotExist)
	//	return
	//}
	result, err := logic.Group.Account.GetAccountByID(c, params.AccountID)
	if err != nil {
		res.Reply(err)
	}
	res.Reply(nil, result)
}

// GetAccountsByName
// @Tags     account
// @Summary  通过昵称模糊查找账户
// @accept   application/form-data
// @Produce  application/json
// @Param    Authorization  header    string            false "x-token 用户令牌"
// @Param    data           query     request.GetAccountsByName                   true  "账号信息"
// @Success  200            {object}  common.State{data=reply.GetAccountsByName}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在"
// @Router   /api/account/infos/name [get]
func (account) GetAccountsByName(c *gin.Context) {
	res := app.NewResponse(c)
	params := &request.GetAccountsByName{}
	if err := c.ShouldBindQuery(params); err != nil {
		zap.S().Errorf("&request.GetAccountByName{} c.ShouldBindJSON(params) failed: %v", err)
		res.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	limit, offset := global.Pager.GetPageSizeAndOffset(c)
	result, err := logic.Group.Account.GetAccountsByName(c, params.AccountName, limit, offset)

	res.ReplyList(err, result.Total, result.AccountInfos)
}

// GetAccountsByUserID
// @Tags     account
// @Summary  获取用户的所有账户
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string          false "x-token 用户令牌"
// @Success  200            {object}  common.State{data=reply.GetAccountsByUserID}  "1003:系统错误 2008:身份验证失败 2010:账号不存在"
// @Router   /api/account/infos/user [get]
func (account) GetAccountsByUserID(c *gin.Context) {
	res := app.NewResponse(c)
	content, ok := middleware.GetPayLoad(c)
	if !ok || content.Type != model.UserToken {
		res.Reply(errcode.AuthNotExist)
		return
	}
	result, err := logic.Group.Account.GetAccountsByUserID(c, content.ID)
	res.ReplyList(err, result.Total, result.AccountInfos)
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
	err := logic.Group.Account.UpdateAccount(c, params.AccountID, params.Name, params.Signature, params.Avatar, params.Gender)
	res.Reply(err)
}
