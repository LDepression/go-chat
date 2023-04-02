package routing

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat/internal/api/v1"
	"go-chat/internal/middleware"
)

type account struct {
}

func (a *account) Init(root *gin.RouterGroup) {
	account := v1.NewAccount()

	ag := root.Group("account").Use(middleware.Auth())
	{
		ag.GET("/info", account.GetAccountByID)
		ag.GET("/infos/name", account.GetAccountsByName)
		ag.PUT("update", account.UpdateAccount)
	}

}
