/**
 * @Author: lenovo
 * @Description:
 * @File:  file
 * @Version: 1.0.0
 * @Date: 2023/04/11 22:36
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat/internal/api/v1"
	"go-chat/internal/middleware"
)

type file struct{}

func (file) Init(root *gin.RouterGroup) {
	//
	rg := root.Group("file").Use(middleware.Auth(), middleware.AuthMustAccount())
	rg.POST("/uploadAccountAvatar", v1.Group.File.UploadAvatar)
	rg.GET("/fileDetails", v1.Group.File.FindFileDetailsByFileID)

}
