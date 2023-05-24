package routing

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat/internal/api/v1"
	"go-chat/internal/middleware"
)

type notify struct {
}

func (n *notify) Init(root *gin.RouterGroup) {
	rg := root.Group("notify").Use(middleware.Auth(), middleware.AuthMustAccount())
	{
		rg.POST("/create", v1.Group.Notify.CreateOneNotify)
		rg.DELETE("/delete", v1.Group.Notify.DeleteNotify)
		rg.GET("/getNotifies", v1.Group.Notify.GetNotifiesByID)
		rg.PUT("/update", v1.Group.Notify.UpdateNotify)
	}
}
