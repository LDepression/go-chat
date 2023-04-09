/**
 * @Author: lenovo
 * @Description:
 * @File:  application
 * @Version: 1.0.0
 * @Date: 2023/04/04 8:30
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

type application struct{}

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
