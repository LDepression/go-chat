/**
 * @Author: lyc
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/03/20 18:41
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat/internal/api/v1"
	"go-chat/internal/middleware"
)

type user struct {
}

func (u *user) Init(root *gin.RouterGroup) {
	root.POST("/register", v1.Group.User.Register)
	root.POST("/login", v1.Group.User.Login)
	g := root.Group("user").Use(middleware.Auth())
	g.POST("/modifyPassword", v1.Group.User.ModifyPassword)
	g.GET("/logout", v1.Group.User.Logout)
}
