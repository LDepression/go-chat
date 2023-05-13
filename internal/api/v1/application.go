package v1

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/api/base"
	"go-chat/internal/global"
	"go-chat/internal/logic"
	"go-chat/internal/middleware"
	"go-chat/internal/model"
	"go-chat/internal/model/request"
	"go-chat/internal/pkg/app"
	"go-chat/internal/pkg/app/errcode"
	"go.uber.org/zap"
)

type application struct{}

// CreateApplication
// @Tags     application
// @Summary  发起好友申请
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string           true "x-token 用户令牌"
// @Param    data           body      request.CreateApplicationReq  true  "发起好友申请"
// @Success  200            {object}  common.State{reply.CreateApplicationRep}             "1001:参数有误 1003:系统错误 2001:鉴权失败 5003:好友已经存在 5004:不能添加自己为好友"
// @Router   /api/v1/application/create [post]
func (a application) CreateApplication(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var req request.CreateApplicationReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.HandleValidatorError(ctx, err)
		zap.S().Infof("ctx.ShouldBindJSON(&req) failed: %v", err)
		return
	}
	content, ok := middleware.GetContent(ctx)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
		return
	}
	replyInfo, err := logic.Group.Application.CreateApplication(ctx, req, uint(content.ID))
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, replyInfo)
}

// DeleteApplication
// @Tags     application
// @Summary  删除已经发送的好友申请
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string           true "x-token 用户令牌"
// @Param    data           body      request.DeleteApplicationReq  true  "删除已经发起的申请"
// @Success  200            {object}  common.State{}             "1001:参数有误 1003:系统错误 2001:鉴权失败 "
// @Router   /api/v1/application/delete [delete]
func (application) DeleteApplication(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var req request.DeleteAccountReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.HandleValidatorError(ctx, err)
		zap.S().Info("ctx.ShouldBindJSON(&req) failed,err", err)
		return
	}
	content, ok := middleware.GetContent(ctx)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
		return
	}
	if err := logic.Group.Application.DeleteApplication(ctx, uint64(req.AccountID)); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

// AcceptApplication
// @Tags     application
// @Summary  被申请者同意好友申请
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string           true "x-token 用户令牌"
// @Param    data           body      request.AcceptApplication  true  "需要同意的申请"
// @Success  200            {object}  common.State{}             "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在 3002:申请不存在 3004:重复操作申请"
// @Router   /api/v1/application/accept [put]
func (application) AcceptApplication(c *gin.Context) {
	res := app.NewResponse(c)
	params := &request.AcceptApplication{}
	if err := c.ShouldBindJSON(params); err != nil {
		zap.S().Infof("&request.AcceptApplication ShouldBindJSON failed,err:%v", err)
		res.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := middleware.GetContent(c)
	if !ok || content.Type != model.AccountToken {
		res.Reply(errcode.AuthNotExist)
		return
	}
	err := logic.Group.Application.AcceptApplication(c, params.ApplicantID, content.ID)
	res.Reply(err)
}

// RefuseApplication
// @Tags     application
// @Summary  被申请者拒绝好友申请
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                     true "x-token 用户令牌"
// @Param    data           body      request.RefuseApplication  true  "需要拒绝的申请"
// @Success  200            {object}  common.State{}             "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在 3002:申请不存在 3004:重复操作申请"
// @Router   /api/v1/application/refuse [put]
func (application) RefuseApplication(c *gin.Context) {
	res := app.NewResponse(c)
	params := &request.RefuseApplication{}
	if err := c.ShouldBindJSON(params); err != nil {
		zap.S().Infof("&request.RefuseApplication ShouldBindJSON failed,err:%v", err)
		res.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := middleware.GetContent(c)
	if !ok || content.Type != model.AccountToken {
		res.Reply(errcode.AuthNotExist)
		return
	}
	err := logic.Group.Application.RefuseApplication(c, params.ApplicantID, content.ID, params.RefuseMsg)
	res.Reply(err)
}

// GetApplicationsList
// @Tags     application
// @Summary  账户查看和自身相关的好友申请(不论是申请者还是被申请者)
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                        true "x-token 用户令牌"
// @Param    data           query     request.GetApplicationsList                  true  "分页参数"
// @Success  200            {object}  common.State{data=reply.ApplicationsList}  "1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在"
// @Router   /api/v1/application/list [get]
func (application) GetApplicationsList(c *gin.Context) {
	res := app.NewResponse(c)
	content, ok := middleware.GetContent(c)
	if !ok || content.Type != model.AccountToken {
		res.Reply(errcode.AuthNotExist)
		return
	}
	limit, offset := global.Pager.GetPageSizeAndOffset(c)
	result, err := logic.Group.Application.GetApplicationsList(c, content.ID, uint(limit), uint(offset))
	res.ReplyList(err, result.Total, result.ApplicationList)
}
