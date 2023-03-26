/**
 * @Author: lenovo
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2023/03/20 18:38
 */

package router

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/middleware"
	"go-chat/internal/routing"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.DebugMode)
	r.Use(middleware.Recovery(true), middleware.LogBody())
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})
	root := r.Group("api/v1")
	routing.Group.User.Init(root)
	routing.Group.Email.Init(root)
	return r
}
