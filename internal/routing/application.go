package routing

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat/internal/api/v1"
	"go-chat/internal/middleware"
)

type application struct{}

func (application) Init(root *gin.RouterGroup) {
	g := root.Group("/application").Use(middleware.Auth(), middleware.AuthMustAccount())
	g.POST("/create", v1.Group.Application.CreateApplication)
	g.DELETE("/delete", v1.Group.Application.DeleteApplication)
	g.PUT("accept", v1.Group.Application.AcceptApplication)
	g.PUT("refuse", v1.Group.Application.RefuseApplication)
	g.GET("list", v1.Group.Application.GetApplicationsList)
}
