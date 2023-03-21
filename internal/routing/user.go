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
)

type user struct {
}

func (u *user) Init(root *gin.RouterGroup) {
	user := v1.NewUser()
	g := root.Group("user")
	g.POST("/register", user.Register)
}
