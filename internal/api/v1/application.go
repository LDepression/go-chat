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

type application struct{}

func NewApplication() *application {
	return &application{}
}

// AcceptApplication
// @Tags     application
// @Summary  被申请者同意好友申请
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string           true "x-token 用户令牌"
// @Param    data           body      request.AcceptApplication  true  "需要同意的申请"
// @Success  200            {object}  common.State{}             "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在 3002:申请不存在 3004:重复操作申请"
// @Router   /api/application/accept [put]
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
// @Router   /api/application/refuse [put]
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
// @Router   /api/application/list [get]
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
