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
	"go-chat/internal/dao/mysql/query"
	"go-chat/internal/global"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app"
	"go-chat/internal/pkg/app/errcode"
	"go.uber.org/zap"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rly := app.NewResponse(ctx)
		tokenString := ctx.GetHeader(global.Settings.Token.AuthType)
		payLoad, err := global.Maker.VerifyToken(tokenString)
		if err != nil {
			rly.Reply(errcode.ErrServer)
			zap.S().Infof("global.Maker.VerifyToken(tokenString) failed: %v", err)
			ctx.Abort()
		}
		userID := payLoad.UserID
		quser := query.NewQueryUser()
		userInfo, err := quser.GetUserByID(userID)
		if err != nil || userInfo.ID == 0 || userInfo.Email == "" {
			rly.Reply(myerr.UserNotFound)
			zap.S().Infof("quser.GetUserByID(userID) failed: %v", err)
			ctx.Abort()
		}
		ctx.Next()
	}
}
