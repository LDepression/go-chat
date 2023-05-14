/**
 * @Author: lenovo
 * @Description:
 * @File:  file
 * @Version: 1.0.0
 * @Date: 2023/04/12 20:37
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

type file struct{}

// UploadAvatar
// @Tags     file
// @Summary  获取账户信息
// @accept   application/form-data
// @Produce  application/json
// @Param    Authorization  header    string               true "x-token 用户令牌"
// @Param    data           query     request.UpdateAvatar                   true  "账号信息"
// @Success  200            {object}  common.State{data=reply.UploadReply}  "1001:参数有误 1003:系统错误 2001:用户不存在"
// @Router   /api/v1/file//uploadAccountAvatar [post]
func (file) UploadAvatar(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var rep request.UpdateAvatar
	if err := ctx.ShouldBind(&rep); err != nil {
		base.HandleValidatorError(ctx, err)
		rly.Reply(errcode.ErrParamsNotValid)
		return
	}
	content, ok := middleware.GetContent(ctx)
	if content.Type != model.AccountToken || !ok {
		rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
		return
	}
	if Avatar, err := logic.Group.File.UpdateAccountAvatar(uint64(content.ID), rep.File); err != nil {
		rly.Reply(errcode.ErrServer)
		zap.S().Info("logic.Group.File.UpdateAccountAvatar(int64(content.ID), rep.File)", zap.Any("err", err))
		return
	} else {
		rly.Reply(nil, Avatar)
	}

}

// FindFileDetailsByFileID
// @Tags     file
// @Summary  获取账户信息
// @accept   application/form-data
// @Produce  application/json
// @Param    Authorization  header    string               true "x-token 用户令牌"
// @Param    data           query     request.FindByFileID                   true  "账号信息"
// @Success  200            {object}  common.State{data=reply.FileDetails}  "1001:参数有误 1003:系统错误 2009:权限不足 2007:身份不存在 2008:身份验证失败 2010:账号不存在"
// @Router   /api/v1/file/fileDetails [get]
func (file) FindFileDetailsByFileID(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var rep request.FindByFileID
	if err := ctx.ShouldBindJSON(&rep); err != nil {
		rly.Reply(errcode.ErrServer)
		base.HandleValidatorError(ctx, err)
		return
	}
	fileDetail, err := logic.Group.File.FindFileInfoByFileID(rep.FileID)
	if err != nil {
		zap.S().Info("logic.Group.File.FindFileInfoByFileID failed: %v", err)
		rly.Reply(err)
		return
	}
	rly.Reply(nil, fileDetail)
}

func (file) FindFilesByRelation(ctx *gin.Context) {

}
