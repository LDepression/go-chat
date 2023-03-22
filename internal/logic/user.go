/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/03/20 20:07
 */

package logic

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/dao"
	"go-chat/internal/dao/mysql/query"
	"go-chat/internal/global"
	"go-chat/internal/model/automigrate"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app/errcode"
	password2 "go-chat/internal/pkg/password"
	"go-chat/internal/pkg/token"
	workemail "go-chat/internal/work/email"
	"go.uber.org/zap"
	"time"
)

type user struct {
}

type tokenResult struct {
	token   string
	PayLoad *token.Payload
	Err     error
}

func CreateToken(resultChan chan<- tokenResult, userID int64, ExpireTime time.Duration) func() {
	return func() {
		defer close(resultChan)
		Token, payLoad, err := global.Maker.CreateToken(userID, ExpireTime)
		resultChan <- tokenResult{
			token:   Token,
			PayLoad: payLoad,
			Err:     err,
		}
	}

}
func (user) Register(ctx *gin.Context, mobile, email, password, code string) errcode.Err {
	//先判断一下邮箱是否已经存在
	exist, err := Group.Email.CheckEmailIsUsed(ctx, email)
	if err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	if exist {
		return myerr.EmailExists
	}
	//再去验证一下邮箱验证码
	if ok := workemail.CheckEmailAndCodeValid(email, code); !ok {
		return myerr.EmailCodeInvalid
	}
	//将验证码存入redis和mysql中去
	if err := dao.Group.Redis.SaveEmail(ctx, email); err != nil {
		zap.S().Infof("dao.Group.Redis.SaveEmail failed ,err:%v", err)
		reTry("addEmail"+email, func() error { return dao.Group.Redis.SaveEmail(ctx, email) })
	}
	var user automigrate.User
	//将密码进行加密
	hashPassword, err := password2.HashPassword(password)
	if err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	user.Password = hashPassword
	user.Mobile = mobile
	user.Email = email
	quser := query.NewQueryUser()
	if err := quser.SaveRegisterInfo(user); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}

	//生成AUTH Token
	accessChan := make(chan tokenResult, 1)
	refreshChan := make(chan tokenResult, 1)
	global.Worker.SendTask(CreateToken(accessChan))
	return nil
}
