/**
 * @Author: lenovo
 * @Description:
 * @File:  email
 * @Version: 1.0.0
 * @Date: 2023/03/20 22:50
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat/internal/api/v1"
)

type email struct{}

func (email) Init(root *gin.RouterGroup) {
	e := root.Group("email")
	e.POST("/sendEmail", v1.Group.Email.SendEmail)
	e.POST("/check", v1.Group.Email.ExistEmail)
}
