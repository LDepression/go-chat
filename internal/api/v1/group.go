/**
 * @Author: lenovo
 * @Description:
 * @File:  group
 * @Version: 1.0.0
 * @Date: 2023/05/06 22:49
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/api/base"
	"go-chat/internal/logic"
	"go-chat/internal/middleware"
	"go-chat/internal/model/request"
	"go-chat/internal/pkg/app"
	"go-chat/internal/pkg/app/errcode"
	"go.uber.org/zap"
)

type vgroup struct{}

func (vgroup) CreateGroup(ctx *gin.Context) {
	var req request.CreateGroupReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	rly := app.NewResponse(ctx)
	content, ok := middleware.GetContent(ctx)
	if !ok {
		zap.S().Info("middleware.GetContent(ctx) err")
		rly.Reply(errcode.AuthNotExist)
		return
	}
	result, err := logic.Group.VGroup.CreateGroup(content.ID, req)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, result)
}

func (vgroup) DissolveGroup(ctx *gin.Context) {
	var req request.DissolveReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	rly := app.NewResponse(ctx)
	content, ok := middleware.GetContent(ctx)
	if !ok {
		zap.S().Info("middleware.GetContent(ctx) err")
		rly.Reply(errcode.AuthNotExist)
		return
	}
	if err := logic.Group.VGroup.Dissolve(ctx, content.ID, req.RelationID); err != nil {
		zap.S().Infof("logic.Group.VGroup.Dissolve failed,err:%v", err)
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

func (vgroup) Invite2Group(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var req request.InviteParamReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zap.S().Info("ctx.shouldBindJSON failed err:", err)
		base.HandleValidatorError(ctx, err)
		return
	}
	content, ok := middleware.GetContent(ctx)
	if !ok {
		zap.S().Info("middleware.GetContent(ctx) err")
		rly.Reply(errcode.AuthNotExist)
		return
	}
	if err := logic.Group.VGroup.Invite2Group(ctx, int64(content.ID), req); err != nil {
		zap.S().Info("logic.Group.VGroup.Invite2Group failed err:", err)
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

func (vgroup) GetGroupList(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	content, ok := middleware.GetContent(ctx)
	if !ok {
		zap.S().Info("middleware.GetContent failed err:")
		rly.Reply(errcode.AuthNotExist)
		return
	}
	resultInfo, err := logic.Group.VGroup.GetGroupList(int64(content.ID))
	if err != nil {
		rly.Reply(err)
		zap.S().Info("logic.Group.VGroup.GetGroupList failed")
		return
	}
	rly.ReplyList(nil, resultInfo.Total, resultInfo.GroupItems)
}

func (vgroup) TransferGroup(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var req request.TransferReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zap.S().Info("ctx.ShouldBindJSON failed err:", err)
		base.HandleValidatorError(ctx, err)
		return
	}
	content, ok := middleware.GetContent(ctx)
	if !ok {
		zap.S().Info("middleware.GetContent failed err:")
		rly.Reply(errcode.AuthNotExist)
		return
	}
	if err := logic.Group.VGroup.TransferGroup(ctx, int64(content.ID), req.RelationID, req.ToID); err != nil {
		zap.S().Infof("logic.Group.VGroup.TransferGroup failed err:%v", err)
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

func (vgroup) QuitGroup(ctx *gin.Context) {
	//********************主动退********************
	rly := app.NewResponse(ctx)
	var req request.QuitReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zap.S().Info("ctx.ShouldBindJSON failed err:", err)
		base.HandleValidatorError(ctx, err)
		return
	}
	content, ok := middleware.GetContent(ctx)
	if !ok {
		zap.S().Info("middleware.GetContent failed err:")
		rly.Reply(errcode.AuthNotExist)
		return
	}
	if err := logic.Group.VGroup.QuitGroup(ctx, int64(content.ID), req.RelationID); err != nil {
		zap.S().Infof("logic.Group.VGroup.TransferGroup failed err:%v", err)
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

func (vgroup) GetGroupMembers(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var req request.GetMems
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zap.S().Infof("ctx.ShouldBindJSON failed err:%v", err)
		base.HandleValidatorError(ctx, err)
		return
	}
	membersInfo, err := logic.Group.VGroup.GetMembers(req.RelationID)
	if err != nil {
		zap.S().Infof("logic.Group.VGroup.VGroup failed err:%v", err)
		rly.Reply(err)
		return
	}
	rly.ReplyList(nil, membersInfo.Total, membersInfo.MembersInfo)
}
