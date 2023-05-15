/**
 * @Author: lenovo
 * @Description:
 * @File:  setting
 * @Version: 1.0.0
 * @Date: 2023/04/18 20:35
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
	"go.uber.org/zap"
)

type setting struct{}

func (s *setting) GetShowsList(ctx *gin.Context) {
	rly := app.NewResponse(ctx)

	content, ok := middleware.GetContent(ctx)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
		return
	}
	repInfo, err := logic.Group.Setting.GetShowsOrderByShowTime(uint64(content.ID))
	if err != nil {
		rly.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}
	rly.ReplyList(nil, repInfo.Total, repInfo.Data)
}

func (s *setting) GetPinsList(ctx *gin.Context) {
	rly := app.NewResponse(ctx)

	content, ok := middleware.GetContent(ctx)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
		return
	}
	repInfo, err := logic.Group.Setting.GetPins(uint64(content.ID))
	if err != nil {
		rly.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}
	rly.ReplyList(nil, repInfo.Total, repInfo.Data)
}
func (s *setting) UpdatePinsInfo(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var req request.UpdatePinsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.HandleValidatorError(ctx, err)
		rly.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}
	content, ok := middleware.GetContent(ctx)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
		return
	}
	if err := logic.Group.Setting.UpdatePins(ctx, content.ID, req.IsPin, req.RelationID); err != nil {
		zap.S().Infof("UpdatePins failed,err: %v", err)
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

func (s *setting) UpdateNickName(ctx *gin.Context) {
	var req request.UpdateNickName
	rly := app.NewResponse(ctx)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	content, ok := middleware.GetContent(ctx)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
		return
	}
	if err := logic.Group.Setting.UpdateNickName(ctx, content.ID, req); err != nil {
		zap.S().Infof("UpdateNickName failed,err: %v", err)
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

func (s *setting) UpdateDisturbState(ctx *gin.Context) {
	rly := app.NewResponse(ctx)

	var req request.UpdateIsDisturbState
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	content, ok := middleware.GetContent(ctx)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
		return
	}
	if err := logic.Group.Setting.UpdateDisturbState(ctx, content.ID, req); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

func (s *setting) UpdateShowState(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var req request.UpdateShowState
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	content, ok := middleware.GetContent(ctx)
	if !ok || content.Type != model.AccountToken {
		rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
		return
	}
	if err := logic.Group.Setting.UpdateShowState(ctx, content.ID, req); err != nil {
		zap.S().Info("logic.Group.Setting.UpdateShowState(ctx, req) err:", err)
		rly.Reply(err)
		return
	}
}
