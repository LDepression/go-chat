/**
 * @Author: lenovo
 * @Description:
 * @File:  email
 * @Version: 1.0.0
 * @Date: 2023/03/20 20:46
 */

package logic

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/dao"
	"go-chat/internal/dao/mysql/query"
	"go-chat/internal/global"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app/errcode"
	email2 "go-chat/internal/work/email"
	"go.uber.org/zap"
)

type email struct {
}

// CheckEmail 先去redis中查询,再去数据库中查询
func (email) CheckEmail(ctx *gin.Context, emailStr string) (bool, error) {
	exists, err := dao.Group.Redis.CheckEmailExists(ctx, emailStr)
	if err != nil {
		return exists, err
	} else {
		if exists {
			return exists, nil
		}
	}
	quser := query.NewQueryUser()
	used, err := quser.CheckEmailBeUsed(emailStr)
	return used, err
}
func (email) CheckEmailIsUsed(ctx *gin.Context, emailStr string) (bool, error) {
	used, err := email{}.CheckEmail(ctx, emailStr)
	if err != nil {
		return false, err
	} else {
		return used, nil
	}
}

func (email) SendEmail(emailStr string) errcode.Err {
	sendEmail := email2.NewSendCodeTask(emailStr)
	if ok := email2.CheckEmailBeMask(emailStr); ok {
		return myerr.EmailSendTooMany
	}
	//异步发送通知
	global.Worker.SendTask(sendEmail.SendTask())
	go func() {
		//这里没有取到值的话,会一直阻塞
		result := sendEmail.GetChanResult()
		if result.Error != nil {
			zap.S().Info(result.Code)
			zap.S().Info(result.Error)
			switch result.Error {
			case email2.ErrSendTooMany:
				zap.S().Info("send too many")
			default:
				zap.S().Infof("send error: %v", result.Error)
			}
		}
	}()
	return nil
}
