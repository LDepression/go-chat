/**
 * @Author: lenovo
 * @Description:
 * @File:  account
 * @Version: 1.0.0
 * @Date: 2023/03/28 17:27
 */

package logic

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-chat/internal/dao/mysql/query"
	tx2 "go-chat/internal/dao/mysql/tx"
	"go-chat/internal/global"
	"go-chat/internal/middleware"
	"go-chat/internal/model"
	"go-chat/internal/model/common"
	reply2 "go-chat/internal/model/reply"
	"go-chat/internal/model/request"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app/errcode"
	"gorm.io/gorm"
)

type account struct{}

func (account) CreateAccount(ctx *gin.Context, req request.CreateAccountReq) (*reply2.CreateAccountReply, errcode.Err) {
	//先来判断一下账户创建的账户是否超过最大了
	content, ok := middleware.GetContent(ctx)
	if !ok || content.Type != model.UserToken {
		return nil, myerr.TokenNotFound
	}
	tx := tx2.NewAccountTX()
	id := global.SnowFlake.GetId()
	req.ID = id
	if err := tx.CreateAccountWithTX(ctx, content.ID, req); err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	//生成token
	accessChan := make(chan tokenResult, 1)
	global.Worker.SendTask(CreateToken(accessChan, model.AccountToken, id, global.Settings.Token.AccountTokenExpire))

	var reply reply2.CreateAccountReply
	reply.AccountID = id
	reply.Name = req.Name
	switch *req.Gender {
	case 0:
		reply.Gender = "男"
	case 1:
		reply.Gender = "女"
	default:
		return nil, errcode.ErrParamsNotValid
	}
	resChan := <-accessChan
	reply.Token = common.Token{
		Token:     resChan.token,
		ExpiresAt: resChan.PayLoad.ExpiredAt,
	}
	return &reply, nil
}

func (account) DeleteAccount(ctx *gin.Context, accountID int64) errcode.Err {
	content, ok := middleware.GetContent(ctx)
	if !ok || content.Type != model.UserToken {
		return myerr.TokenNotFound
	}
	tx := tx2.NewAccountTX()
	if err := tx.DeleteAccountWithTX(ctx, accountID); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}

func (account) GetToken(ctx *gin.Context, accountID int64) (common.Token, errcode.Err) {
	var replyInfo common.Token
	content, ok := middleware.GetContent(ctx)
	//先去检验一下
	if !ok || content.Type != model.UserToken {
		return replyInfo, myerr.TokenNotFound
	}
	//先去判断一下,id是否是存在的
	qAccount := query.NewAccount()
	accountInfo, err := qAccount.CheckAccountInfoByAccountID(accountID)
	if err != nil || accountInfo.ID == 0 {
		return replyInfo, myerr.AccountNotExist
	}
	//再去数据库里面检验一下,防止被造爬虫数据
	if accountInfo.UserID != uint(content.ID) {
		return common.Token{}, errcode.ErrInsufficientPermissions
	}
	accessChan := make(chan tokenResult, 1)
	global.Worker.SendTask(CreateToken(accessChan, model.AccountToken, accountID, global.Settings.Token.AccountTokenExpire))
	accessRes := <-accessChan
	return common.Token{
		Token:     accessRes.token,
		ExpiresAt: accessRes.PayLoad.ExpiredAt,
	}, nil
}

func (account) GetAccountsByUserID(userID int64) (reply2.TotalAccountsReply, errcode.Err) {
	var reply reply2.TotalAccountsReply

	//先去查询一下userID是否是存在的]
	qUser := query.NewQueryUser()
	userInfo, err := qUser.GetUserByID(userID)
	if err != nil {
		return reply, errcode.ErrServer
	}
	if userInfo.ID == 0 {
		return reply, myerr.UserNotExist
	}
	qAccount := query.NewAccount()
	AccountsInfo, err := qAccount.GetAccountsByUserID(userID)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return reply, nil
		}
		return reply, errcode.ErrServer.WithDetails(err.Error())
	}
	reply.Total = len(AccountsInfo)
	for _, accountInfo := range AccountsInfo {
		var accountInfoReply reply2.AccountInfoReply
		accountInfoReply.AccountID = int64(accountInfo.ID)
		accountInfoReply.Name = accountInfo.Name
		accountInfoReply.Signature = accountInfo.Signature
		accountInfoReply.Avatar = accountInfo.Signature
		reply.AccountInfos = append(reply.AccountInfos, accountInfoReply)
	}
	return reply, nil
}
