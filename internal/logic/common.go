/**
 * @Author: lenovo
 * @Description:
 * @File:  common
 * @Version: 1.0.0
 * @Date: 2023/03/21 21:57
 */

package logic

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/global"
	"go-chat/internal/pkg/retry"
	"go-chat/internal/pkg/token"
)

// 尝试重试
// 失败: 打印日志
func reTry(name string, f func() error) {
	go func() {
		d := global.Settings.Auto.Retry.TimeDuration
		times := global.Settings.Auto.Retry.TimeCount
		report := <-retry.NewTry(name, f, d, times).Run()
		global.Logger.Error(report.Error())
	}()
}

func GetPayLoad(ctx *gin.Context) (string, *token.Payload, error) {
	tokenString := ctx.GetHeader(global.Settings.Token.AuthType)
	payLoad, err := global.Maker.VerifyToken(tokenString)
	if err != nil {
		return "", nil, err
	}
	return tokenString, payLoad, nil
}
