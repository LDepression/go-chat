package routing

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat/internal/api/v1"
	"go-chat/internal/middleware"
)

type setting struct {
}

func (s *setting) Init(root *gin.RouterGroup) {
	g := root.Group("setting", middleware.Auth())
	{
		friendGroup := g.Group("friend", middleware.AuthMustAccount())
		{
			friendGroup.DELETE("delete", v1.Group.Setting.DeleteFriend)
			friendGroup.GET("list", v1.Group.Setting.GetFriendsList)
			friendGroup.GET("list/name", v1.Group.Setting.GetFriendsByName)

			friendGroup.PUT("update/nick_name", v1.Group.Setting.UpdateNickName)
			friendGroup.PUT("update/disturb", v1.Group.Setting.UpdateSettingDisturb)
		}
	}

}
