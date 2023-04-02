/**
 * @Author: lenovo
 * @Description:
 * @File:  account
 * @Version: 1.0.0
 * @Date: 2023/03/28 16:33
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat/internal/api/v1"
	"go-chat/internal/middleware"
)

type account struct{}

func (account) Init(root *gin.RouterGroup) {
	g := root.Group("account").Use(middleware.Auth(), middleware.AuthMustUser())
	g.POST("/createAccount", v1.Group.Account.CreateAccount)
	g.DELETE("/deleteAccount/:id", v1.Group.Account.DeleteAccount)
	g.GET("/getToken/:id", v1.Group.Account.GetToken)
	g.GET("/infos/user", v1.Group.Account.GetAccountsByUserID)
}
