package logic

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/dao/mysql/query"
	"go-chat/internal/model/reply"
	"go-chat/internal/pkg/app/errcode"
	"go.uber.org/zap"
)

type account struct {
}

func (account) GetAccountByID(c *gin.Context, accountID int64) (*reply.GetAccountByID, errcode.Err) {
	qAccount := query.NewQueryAccount()
	accountInfo, err := qAccount.GetAccountByID(accountID)
	if err != nil {
		zap.S().Error("dao.qAccount.GetAccountByID() failed:%v", zap.Error(err))
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	return &reply.GetAccountByID{
		Account: reply.Account{
			ID:        accountInfo.ID,
			CreatedAt: accountInfo.CreatedAt,
			UserID:    accountInfo.UserID,
			Name:      accountInfo.Name,
			Signature: accountInfo.Signature,
			Avatar:    accountInfo.Avatar,
			Gender:    accountInfo.Gender,
		},
	}, nil
}

func (account) GetAccountsByName(c *gin.Context, accountName string, limit, offset int32) (*reply.GetAccountsByName, errcode.Err) {
	qAccount := query.NewQueryAccount()
	accountInfos, totalCount, err := qAccount.GetAccountsByName(accountName, limit, offset)
	if err != nil {
		zap.S().Error("dao.qAccount.GetAccountByName() failed:%v", zap.Error(err))
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	replyAccountInfos := make([]*reply.Account, 0, len(accountInfos))
	for _, v := range accountInfos {
		replyAccountInfos = append(replyAccountInfos, &reply.Account{
			ID:        v.ID,
			CreatedAt: v.CreatedAt,
			UserID:    v.UserID,
			Name:      v.Name,
			Signature: v.Signature,
			Avatar:    v.Avatar,
			Gender:    v.Gender,
		})
	}

	return &reply.GetAccountsByName{
		Account: replyAccountInfos,
		Total:   totalCount,
	}, nil
}
