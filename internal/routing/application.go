package routing

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat/internal/api/v1"
	"go-chat/internal/middleware"
)

type application struct {
}

func (a *application) Init(root *gin.RouterGroup) {
	g := root.Group("application").Use(middleware.Auth(), middleware.AuthMustAccount())
	{
		g.PUT("accept", v1.Group.Application.AcceptApplication)
		g.PUT("refuse", v1.Group.Application.RefuseApplication)
		g.GET("list", v1.Group.Application.GetApplicationsList)
	}
}
