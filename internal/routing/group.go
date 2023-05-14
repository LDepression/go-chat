/**
 * @Author: lenovo
 * @Description:
 * @File:  group
 * @Version: 1.0.0
 * @Date: 2023/05/07 17:35
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat/internal/api/v1"
	"go-chat/internal/middleware"
)

type vgroup struct{}

func (vgroup) Init(r *gin.RouterGroup) {
	rg := r.Group("/group", middleware.Auth(), middleware.AuthMustAccount())
	{
		rg.POST("/create", v1.Group.VGroup.CreateGroup)
		rg.POST("/dissolve", v1.Group.VGroup.DissolveGroup)
		rg.POST("/invite", v1.Group.VGroup.Invite2Group)
		rg.POST("/transfer", v1.Group.VGroup.TransferGroup)
		rg.POST("/quit", v1.Group.VGroup.QuitGroup)
		rg.GET("/list", v1.Group.VGroup.GetGroupList)
		rg.GET("/members", v1.Group.VGroup.GetGroupMembers)
	}
}
