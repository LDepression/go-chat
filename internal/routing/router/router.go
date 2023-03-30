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

	_ "go-chat/docs" // 千万不要忘了导入把你上一步生成的docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.DebugMode)
	r.Use(middleware.Recovery(true), middleware.LogBody(), middleware.Cors())
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	root := r.Group("api/v1")
	routing.Group.User.Init(root)
	routing.Group.Email.Init(root)
	return r
}
