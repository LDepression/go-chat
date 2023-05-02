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

type setting struct {
}

// DeleteFriend
// @Tags     setting
// @Summary  删除好友关系(双向删除)
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                true  "x-token 用户令牌"
// @Param    data           body      request.DeleteFriend  true  "对方账户ID"
// @Success  200            {object}  common.State{}        "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2009:权限不足 2010:账号不存在 4002:关系不存在"
// @Router   /api/setting/friend/delete [delete]
func (setting) DeleteFriend(c *gin.Context) {
	res := app.NewResponse(c)
	params := &request.DeleteFriend{}
	if err := c.ShouldBindJSON(params); err != nil {
		zap.S().Errorf("&request.DeleteFriend{} c.ShouldBindJSON(params) failed: %v", err)
		res.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := middleware.GetContent(c)
	if !ok || content.Type != model.AccountToken {
		res.Reply(errcode.AuthNotExist)
		return
	}
	err := logic.Group.Setting.DeleteFriend(c, content.ID, params.TargetAccountID)
	res.Reply(err)
}

// GetFriendsList
// @Tags     setting
// @Summary  获取当前账户的好友列表
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string      true  "x-token 账户令牌"
// @Success  200            {object}  common.State{data=reply.GetFriendsList}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在"
// @Router   /api/setting/friend/list [get]
func (setting) GetFriendsList(c *gin.Context) {
	res := app.NewResponse(c)
	content, ok := middleware.GetContent(c)
	if !ok || content.Type != model.AccountToken {
		res.Reply(errcode.AuthNotExist)
		return
	}
	result, err := logic.Group.Setting.GetFriendsList(c, content.ID)
	res.ReplyList(err, int64(result.Total), result.FriendsInfos)
}

// GetFriendsByName
// @Tags     setting
// @Summary  通过姓名模糊查询好友(好友姓名或者昵称)
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                                     true  "x-token 账户令牌"
// @Param    data           query     request.GetFriendsByName                   true  "关系ID，免打扰状态"
// @Success  200            {object}  common.State{data=reply.GetFriendsByName}  "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在"
// @Router   /api/setting/friend/name [get]
func (setting) GetFriendsByName(c *gin.Context) {
	res := app.NewResponse(c)
	params := &request.GetFriendsByName{}
	if err := c.ShouldBindJSON(params); err != nil {
		zap.S().Errorf("&request.GetFriendsByName{} c.ShouldBindJSON(params) failed: %v", err)
		res.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := middleware.GetContent(c)
	if !ok || content.Type != model.AccountToken {
		res.Reply(errcode.AuthNotExist)
		return
	}

	limit, offset := global.Pager.GetPageSizeAndOffset(c)
	result, err := logic.Group.Setting.GetFriendsByName(c, content.ID, uint(limit), uint(offset), params.Name)
	res.ReplyList(err, int64(result.Total), result.FriendsInfos)
}

// UpdateNickName
// @Tags     setting
// @Summary  更新好友备注或群昵称
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                  true  "x-token 账户令牌"
// @Param    data           body      request.UpdateNickName  true  "关系ID，备注或群昵称"
// @Success  200            {object}  common.State{}          "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在 4002:关系不存在"
// @Router   /api/setting/update/nick_name [put]
func (setting) UpdateNickName(c *gin.Context) {
	res := app.NewResponse(c)
	params := &request.UpdateNickName{}
	if err := c.ShouldBindJSON(params); err != nil {
		zap.S().Errorf("&request.UpdateNickName{} c.ShouldBindJSON(params) failed: %v", err)
		res.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := middleware.GetContent(c)
	if !ok || content.Type != model.AccountToken {
		res.Reply(errcode.AuthNotExist)
		return
	}
	err := logic.Group.Setting.UpdateNickName(c, content.ID, params.RelationID, params.NickName)
	res.Reply(err)
}

// UpdateSettingDisturb
// @Tags     setting
// @Summary  更改好友或群组免打扰选项
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string                     true  "x-token 账户令牌"
// @Param    data           body      request.UpdateSettingDisturb  true  "关系ID，免打扰状态"
// @Success  200            {object}  common.State{}             "1001:参数有误 1003:系统错误 2007:身份不存在 2008:身份验证失败 2010:账号不存在 4002:关系不存在"
// @Router   /api/setting/update/disturb [put]
func (setting) UpdateSettingDisturb(c *gin.Context) {
	res := app.NewResponse(c)
	params := &request.UpdateSettingDisturb{}
	if err := c.ShouldBindJSON(params); err != nil {
		zap.S().Errorf("&request.UpdateNickName{} c.ShouldBindJSON(params) failed: %v", err)
		res.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := middleware.GetContent(c)
	if !ok || content.Type != model.AccountToken {
		res.Reply(errcode.AuthNotExist)
		return
	}
	err := logic.Group.Setting.UpdateSettingDisturb(c, content.ID, params.RelationID, params.IsDisturbed)
	res.Reply(err)
}
