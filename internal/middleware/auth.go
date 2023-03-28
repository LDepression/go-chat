/**
 * @Author: lenovo
 * @Description:
 * @File:  auth
 * @Version: 1.0.0
 * @Date: 2023/03/25 14:59
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/global"
	"go-chat/internal/model"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app"
	"go-chat/internal/pkg/app/errcode"
	"go.uber.org/zap"
)

func GetPayLoad(ctx *gin.Context) (*model.Content, bool) {
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
		}
		content := &model.Content{}
		err = content.UnMarshal(payLoad.Content)
		if err != nil {
			zap.S().Infof("global.Maker.UnMarshal failed,err: %v", err)
			rly.Reply(errcode.ErrServer)
			ctx.Abort()
		}
		zap.S().Info(content)
		ctx.Set(global.Settings.Token.AuthKey, content)

		ctx.Next()
	}
}
