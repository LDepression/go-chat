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
	"go-chat/internal/middleware"
	"go-chat/internal/model"
	"go-chat/internal/model/automigrate"
	"go-chat/internal/model/reply"
	"go-chat/internal/model/request"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app/errcode"
	password2 "go-chat/internal/pkg/password"
	"go-chat/internal/pkg/token"
	workemail "go-chat/internal/work/email"
	"go.uber.org/zap"
	"strings"
	"time"
)

type user struct {
}

type tokenResult struct {
	token   string
	PayLoad *token.Payload
	Err     error
}

func CreateToken(resultChan chan<- tokenResult, t model.TokenType, id int64, ExpireTime time.Duration) func() {
	return func() {
		defer close(resultChan)
		content, _ := model.NewContent(t, id).Marshal()
		zap.S().Info(string(content))
		Token, payLoad, err := global.Maker.CreateToken(content, ExpireTime)
		resultChan <- tokenResult{
			token:   Token,
			PayLoad: payLoad,
			Err:     err,
		}
	}

}
func (user) Register(ctx *gin.Context, mobile, email, password, code string) (*reply.LoginReply, errcode.Err) {
	//先判断一下邮箱是否已经存在
	exist, err := Group.Email.CheckEmailIsUsed(ctx, email)
	if err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	if exist {
		return nil, myerr.EmailExists
	}
	//再去验证一下邮箱验证码
	if ok := workemail.CheckEmailAndCodeValid(email, code); !ok {
		return nil, myerr.EmailCodeInvalid
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
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	user.Password = hashPassword
	user.Mobile = mobile
	user.Email = email
	quser := query.NewQueryUser()
	userID, err := quser.SaveRegisterInfo(user)
	if err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}

	//生成AUTH Token
	accessChan := make(chan tokenResult, 1)
	refreshChan := make(chan tokenResult, 1)
	global.Worker.SendTask(CreateToken(accessChan, model.UserToken, int64(userID), global.Settings.Token.AccessTokenExpire))
	global.Worker.SendTask(CreateToken(refreshChan, model.UserToken, int64(userID), global.Settings.Token.RefreshTokenExpire))
	accessRes := <-accessChan
	refreshRes := <-refreshChan

	if err := dao.Group.Redis.SaveUserToken(ctx, int64(userID), []string{accessRes.token, refreshRes.token}); err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	return &reply.LoginReply{
		UserID:       int64(userID),
		AccessToken:  accessRes.token,
		RefreshToken: refreshRes.token,
		PayLoad:      *accessRes.PayLoad,
	}, nil
}

func (user) Login(ctx *gin.Context, req request.Login) (*reply.LoginReply, errcode.Err) {

	//先要判断一下用户是否已经登录了

	quser := query.NewQueryUser()
	exist, err := quser.CheckEmailBeUsed(req.Email)
	if err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	if !exist {
		return nil, myerr.EmailNotFound
	}
	userInfo, err := quser.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	count := dao.Group.Redis.CountUserToken(ctx, int64(userInfo.ID))
	if count != 0 {
		return nil, myerr.UserExist
	}
	switch req.LoginType {
	case request.LoginByPassword:
		//先去判断一下密码是否存在
		if strings.TrimSpace(req.Password) == "" {
			return nil, myerr.PasswordInvalid
		}
		if err := password2.CheckPassword(req.Password, userInfo.Password); err != nil {
			return nil, myerr.PasswordError
		}
	case request.LoginByEmail:
		if strings.TrimSpace(req.EmailCode) == "" {
			return nil, myerr.EmailCodeInvalid
		}
		if ok := workemail.CheckEmailAndCodeValid(req.Email, req.EmailCode); !ok {
			return nil, myerr.EmailCodeInvalid
		}
	default:
		return nil, myerr.ChoiceNotFound
	}

	//生成AUTH Token
	accessChan := make(chan tokenResult, 1)
	refreshChan := make(chan tokenResult, 1)
	global.Worker.SendTask(CreateToken(accessChan, model.UserToken, int64(userInfo.ID), global.Settings.Token.AccessTokenExpire))
	global.Worker.SendTask(CreateToken(refreshChan, model.UserToken, int64(userInfo.ID), global.Settings.Token.RefreshTokenExpire))
	accessRes := <-accessChan
	refreshRes := <-refreshChan

	if err := dao.Group.Redis.SaveUserToken(ctx, int64(userInfo.ID), []string{accessRes.token, refreshRes.token}); err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	return &reply.LoginReply{
		UserID:       int64(userInfo.ID),
		AccessToken:  accessRes.token,
		RefreshToken: refreshRes.token,
		PayLoad:      *accessRes.PayLoad,
	}, nil

}

func (user) ModifyPassword(ctx *gin.Context, req request.ReqModifyPassword) errcode.Err {
	Token, payLoad, err := GetPayLoad(ctx)
	if err != nil {
		zap.S().Info("get payLoad failed", zap.Any("error", err.Error()))
		return errcode.ErrServer.WithDetails(err.Error())
	}

	//判断一下token是否失效

	quser := query.NewQueryUser()
	zap.S().Info(payLoad.Content)
	content, exist := middleware.GetPayLoad(ctx)
	if !exist {
		return myerr.TokenNotFound
	}
	userInfo, err := quser.GetUserByID(content.ID)
	if ok := dao.Group.Redis.CheckUserTokenValid(ctx, int64(userInfo.ID), Token); ok == false {
		return myerr.TokenInValid
	}
	//验证emailCode
	if ok := workemail.CheckEmailAndCodeValid(userInfo.Email, req.EmailCode); !ok {
		return myerr.EmailCodeInvalid
	}
	hashPassword, err := password2.HashPassword(req.Password)
	if err != nil {
		zap.S().Info("hashPassword failed", zap.Any("error", err.Error))
		return errcode.ErrServer.WithDetails(err.Error())
	}
	if err = quser.ModifyPassword(userInfo.Email, hashPassword); err != nil {
		zap.S().Info("quser.ModifyPassword failed,", zap.Any("error", err.Error()))
		return errcode.ErrServer.WithDetails(err.Error())
	}

	//现在清除用户的token
	if err := dao.Group.Redis.DeleteAllTokenByUser(ctx, int64(userInfo.ID)); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}

func (user) Logout(ctx *gin.Context) errcode.Err {
	Token, payLoad, err := GetPayLoad(ctx)
	if err != nil {
		return myerr.TokenInValid
	}
	content := &model.Content{}
	_ = content.UnMarshal(payLoad.Content)

	//先判断用户在redis中是否存在
	if ok := dao.Group.Redis.CheckUserTokenValid(ctx, content.ID, Token); !ok {
		return myerr.UserNotExist
	}
	//先将token从redis中清除
	if err := dao.Group.Redis.DeleteAllTokenByUser(ctx, content.ID); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}
