/**
 * @Author: lenovo
 * @Description:
 * @File:  application'
 * @Version: 1.0.0
 * @Date: 2023/04/04 8:29
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat/internal/api/v1"
	"go-chat/internal/middleware"
)

type application struct{}

func (application) Init(root gin.RouterGroup) {
	g := root.Group("/application").Use(middleware.AuthMustAccount())
	g.POST("/createApplication", v1.Group.Application.CreateApplication)
	g.DELETE("/deleteApplication", v1.Group.Application.DeleteApplication)
}
