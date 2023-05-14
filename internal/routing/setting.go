/**
 * @Author: lenovo
 * @Description:
 * @File:  setting
 * @Version: 1.0.0
 * @Date: 2023/04/18 20:30
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat/internal/api/v1"
	"go-chat/internal/middleware"
)

type setting struct{}

func (s *setting) Init(root *gin.RouterGroup) {
	rg := root.Group("setting", middleware.Auth(), middleware.AuthMustAccount())

	{
		rg.GET("/pins", v1.Group.Setting.GetPinsList)
		rg.GET("/shows", v1.Group.Setting.GetShowsList)

		upGroup := rg.Group("update")
		upGroup.PUT("/pins", v1.Group.Setting.UpdatePinsInfo)
		upGroup.PUT("/nick_name", v1.Group.Setting.UpdateNickName)
		upGroup.PUT("/disturb", v1.Group.Setting.UpdateDisturbState)
		upGroup.PUT("/shows", v1.Group.Setting.UpdateShowState)
	}
}
