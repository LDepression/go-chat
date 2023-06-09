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
	socketio "github.com/googollee/go-socket.io"
	"go-chat/internal/middleware"
	"go-chat/internal/routing"
	"net/http"

	_ "go-chat/docs" // 千万不要忘了导入把你上一步生成的docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() (*gin.Engine, *socketio.Server) {

	r := gin.New()
	gin.SetMode(gin.DebugMode)
	gin.ForceConsoleColor()
	root := r.Use(middleware.Recovery(true), middleware.LogBody(), middleware.Cors())
	{
		root.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "pong")
		})
		root.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		r := r.Group("api/v1")
		rg := routing.Group
		rg.User.Init(r)
		rg.Email.Init(r)
		rg.Account.Init(r)
		rg.Application.Init(r)
		rg.File.Init(r)
		rg.Setting.Init(r)
		rg.Group.Init(r)
	}

	return r, routing.Group.Chat.Init(r)
}
