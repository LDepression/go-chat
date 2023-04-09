/**
 * @Author: lenovo
 * @Description:
 * @File:  application
 * @Version: 1.0.0
 * @Date: 2023/04/04 8:51
 */

package logic

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/dao/mysql/query"
	tx2 "go-chat/internal/dao/mysql/tx"
	reply2 "go-chat/internal/model/reply"
	"go-chat/internal/model/request"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app/errcode"
)

type application struct{}

func (application) CreateApplication(ctx *gin.Context, req request.CreateApplicationReq, accountID uint) (*reply2.CreateApplicationRep, errcode.Err) {
	var rep = new(reply2.CreateApplicationRep)
	tx := tx2.NewApplicationTX()
	applicationID, err := tx.CreateApplicationWithTX(uint64(accountID), req.AccountID, req.ApplicationMsg)
	if err != nil {
		switch err {
		case tx2.ErrHasThisFriend:
			return nil, myerr.FriendHasAlreadyExists
		case tx2.ErrIsSelf:
			return nil, myerr.CanNotAddSelf
		default:
			return nil, errcode.ErrServer
		}
	}
	rep.ApplicationID = applicationID
	rep.Status = string(query.WAITING)

	//TODO:提示对方有新的消息
	return rep, nil
}

func (application) DeleteApplication(ctx *gin.Context, accountID uint64) errcode.Err {
	//
	q := query.NewApplication()
	if err := q.DeleteApplication(accountID); err != nil {
		return errcode.ErrServer
	}
	return nil
}
