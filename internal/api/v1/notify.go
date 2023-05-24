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
	"go.uber.org/zap"
)

type notify struct {
}

// CreateOneNotify
// @Tags     notify
// @Summary  创建群通知
// @accept   notify/json
// @Produce  notify/json
// @Param    Authorization  header    string                               true "x_token 用户令牌"
// @Param    data           body      request.CreateNotify                  true  "请求信息"
// @Success  200            {object}  common.State{data=reply.Notify}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 7001:非群主 7003:非群成员"
// @Router   /api/v1/notify/create [post]
func (notify) CreateOneNotify(c *gin.Context) {
	res := app.NewResponse(c)
	params := &request.CreateNotify{}
	if err := c.ShouldBindJSON(params); err != nil {
		zap.S().Errorf("&request.CreateNotify{} c.ShouldBindJSON(params) failed: %v", err)
		base.HandleValidatorError(c, err)
		return
	}
	content, ok := middleware.GetContent(c)
	if !ok || content.Type != model.AccountToken {
		res.Reply(errcode.AuthNotExist)
		return
	}
	result, err := logic.Group.Notify.CreateNotify(c, content.ID, params)
	res.Reply(err, result)
}

// DeleteNotify
// @Tags     notify
// @Summary  删除群通知
// @accept   notify/json
// @Produce  notify/json
// @Param    Authorization  header    string                                true  "x_token 用户令牌"
// @Param    data           query     request.DeleteNotify                 true  "请求信息"
// @Success  200            {object}  common.State{}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 7001:非群主 7003:非群成员"
// @Router   /api/v1/notify/delete [delete]
func (notify) DeleteNotify(c *gin.Context) {
	res := app.NewResponse(c)
	params := &request.DeleteNotify{}
	if err := c.ShouldBindJSON(params); err != nil {
		zap.S().Errorf("&request.DeleteNotify{} c.ShouldBindJSON(params) failed: %v", err)
		base.HandleValidatorError(c, err)
		return
	}
	content, ok := middleware.GetContent(c)
	if !ok || content.Type != model.AccountToken {
		res.Reply(errcode.AuthNotExist)
		return
	}
	err := logic.Group.Notify.DeleteNotify(c, content.ID, params)
	res.Reply(err, nil)
}

// GetNotifiesByID
// @Tags     notify
// @Summary  获取群通知
// @accept   notify/json
// @Produce  notify/json
// @Param    Authorization  header    string                                true  "x_token 用户令牌"
// @Param    data           query     request.GetNotifyByID                 true  "请求信息"
// @Success  200            {object}  common.State{data=reply.GetNotifies}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 7003:非群成员"
// @Router   /api/v1/notify/getNotifies [get]
func (notify) GetNotifiesByID(c *gin.Context) {
	res := app.NewResponse(c)
	params := &request.GetNotifyByID{}
	if err := c.ShouldBindJSON(params); err != nil {
		zap.S().Errorf("&request.GetNotifyByID{} c.ShouldBindJSON(params) failed: %v", err)
		base.HandleValidatorError(c, err)
		return
	}
	content, ok := middleware.GetContent(c)
	if !ok || content.Type != model.AccountToken {
		res.Reply(errcode.AuthNotExist)
		return
	}
	result, err := logic.Group.Notify.GetNotifiesByID(c, content.ID, params)
	res.Reply(err, result)
}

// UpdateNotify
// @Tags     notify
// @Summary  更新群通知
// @accept   notify/json
// @Produce  notify/json
// @Param    Authorization  header    string                                 true  "x_token 用户令牌"
// @Param    data           body      request.UpdateNotify                   true  "请求信息"
// @Success  200            {object}  common.State{data=reply.UpdateNotify}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 7003:非群成员"
// @Router   /api/v1/notify/update [put]
func (notify) UpdateNotify(c *gin.Context) {
	res := app.NewResponse(c)
	params := &request.UpdateNotify{}
	if err := c.ShouldBindJSON(params); err != nil {
		zap.S().Errorf("&request.UpdateNotify{} c.ShouldBindJSON(params) failed: %v", err)
		base.HandleValidatorError(c, err)
		return
	}
	content, ok := middleware.GetContent(c)
	if !ok || content.Type != model.AccountToken {
		res.Reply(errcode.AuthNotExist)
		return
	}
	result, err := logic.Group.Notify.UpdateNotify(c, content.ID, params)
	res.Reply(err, result)
}
