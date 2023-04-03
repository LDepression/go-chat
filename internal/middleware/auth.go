/**
 * @Author: lenovo
 * @Description:
 * @File:  auth
 * @Version: 1.0.0
 * @Date: 2023/03/25 14:59
 */

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-chat/internal/dao/mysql/query"
	"go-chat/internal/global"
	"go-chat/internal/model"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app"
	"go-chat/internal/pkg/app/errcode"
	"go.uber.org/zap"
)

func GetContent(ctx *gin.Context) (*model.Content, bool) {
	Content, ok := ctx.Get(global.Settings.Token.AuthKey)
	c := Content.(*model.Content)
	return c, ok
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rly := app.NewResponse(ctx)
		tokenString := ctx.GetHeader(global.Settings.Token.AuthType)
		payLoad, err := global.Maker.VerifyToken(tokenString)
		if err != nil {
			rly.Reply(myerr.TokenInValid)
			zap.S().Infof("global.Maker.VerifyToken(tokenString) failed: %v", err)
			ctx.Abort()
			return
		}
		content := &model.Content{}
		err = content.UnMarshal(payLoad.Content)
		if err != nil {
			zap.S().Infof("global.Maker.UnMarshal failed,err: %v", err)
			rly.Reply(errcode.ErrServer)
			ctx.Abort()
			return
		}
		zap.S().Info(content)
		ctx.Set(global.Settings.Token.AuthKey, content)

		ctx.Next()
	}
}

func AuthMustUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rly := app.NewResponse(ctx)
		content, exist := GetContent(ctx)
		if !exist {
			rly.Reply(myerr.TokenNotFound)
			ctx.Abort()
			return
		}
		if content.Type != model.UserToken {
			rly.Reply(myerr.AuthFailed)
			ctx.Abort()
			return
		}
		quser := query.NewQueryUser()
		userInfo, err := quser.GetUserByID(content.ID)
		if err != nil {
			rly.Reply(errcode.ErrServer.WithDetails(err.Error()))
			ctx.Abort()
			return
		}
		if userInfo == nil {
			rly.Reply(myerr.UserNotExist)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func AuthMustAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := app.NewResponse(ctx)
		content, exist := GetContent(ctx)
		if !exist {
			res.Reply(myerr.TokenNotFound)
			fmt.Println("1")
			ctx.Abort()
			return
		}
		if content.Type != model.AccountToken {
			res.Reply(myerr.AuthFailed)
			fmt.Println("2")
			ctx.Abort()
			return
		}
		qAccount := query.NewQueryAccount()
		accountInfo, err := qAccount.GetAccountByID(content.ID)
		if err != nil {
			res.Reply(errcode.ErrServer.WithDetails(err.Error()))
			fmt.Println("3")
			ctx.Abort()
			return
		}
		if accountInfo == nil {
			res.Reply(myerr.AccountNotExist)
			fmt.Println("4")
			ctx.Abort()
			return
		}
	}
}
