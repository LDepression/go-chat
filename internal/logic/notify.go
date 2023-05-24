package logic

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/dao/mysql/query"
	"go-chat/internal/model/reply"
	"go-chat/internal/model/request"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app/errcode"
	"go.uber.org/zap"
)

type notify struct {
}

func (notify) CreateNotify(c *gin.Context, accountID uint, params *request.CreateNotify) (*reply.Notify, errcode.Err) {
	qNotify := query.NewQueryNotify()
	IsLeader, err := qNotify.CheckIsLeader(accountID, params.RelationID)
	if err != nil {
		zap.S().Errorf("CreateNotify: qNotify.CheckIsLeader failes,err:%v", err)
		return nil, errcode.ErrServer
	}
	if !IsLeader {
		return nil, myerr.DoNotHaveAuth
	}
	notify, err := qNotify.CreateNotify(accountID, params.RelationID, params.MsgContent, params.MsgExpand)
	if err != nil {
		zap.S().Errorf("CreateNotify: qNotify.CreateNotify failes,err:%v", err)
		return nil, errcode.ErrServer
	}
	replyNotify := &reply.Notify{
		ID:         notify.ID,
		RelationID: notify.RelationID,
		MsgContent: notify.MsgContent,
		MsgExpand:  notify.MsgExpand,
		AccountID:  notify.AccountID,
		UpdateAt:   notify.UpdatedAt,
	}
	//TODO 向群成员通知 msgNotify
	return replyNotify, nil
}

func (notify) DeleteNotify(c *gin.Context, accountID uint, params *request.DeleteNotify) errcode.Err {
	qNotify := query.NewQueryNotify()
	IsLeader, err := qNotify.CheckIsLeader(accountID, params.RelationID)
	if err != nil {
		zap.S().Errorf("CreateNotify: qNotify.CheckIsLeader failes,err:%v", err)
		return errcode.ErrServer
	}
	if !IsLeader {
		return myerr.DoNotHaveAuth
	}
	err = qNotify.DeleteNotify(params.ID, params.RelationID)
	if err != nil {
		zap.S().Errorf("DeleteNotify: qNotify.DeleteNotify failes,err:%v", err)
		return errcode.ErrServer
	}
	return nil
}

func (notify) GetNotifiesByID(c *gin.Context, accountID uint, params *request.GetNotifyByID) (*reply.GetNotifies, errcode.Err) {
	qNotify := query.NewQueryNotify()
	IsInGroup, err := qNotify.CheckIsInGroup(accountID, params.RelationID)
	if err != nil {
		zap.S().Errorf("CreateNotify: qNotify.CheckIsLeader failes,err:%v", err)
		return &reply.GetNotifies{}, errcode.ErrServer
	}
	if !IsInGroup {
		return &reply.GetNotifies{}, myerr.DoNotHaveThisAccount
	}
	notifiesInfo, err := qNotify.GetNotifies(params.RelationID)
	if err != nil {
		zap.S().Errorf("CreateNotify: qNotify.GetNotifies failes,err:%v", err)
		return &reply.GetNotifies{}, errcode.ErrServer
	}
	replyNotifiesInfo := make([]reply.Notify, 0)
	for _, v := range notifiesInfo {
		replyNotifiesInfo = append(replyNotifiesInfo, reply.Notify{
			ID:         v.ID,
			RelationID: v.RelationID,
			MsgContent: v.MsgContent,
			MsgExpand:  v.MsgExpand,
			AccountID:  v.AccountID,
			UpdateAt:   v.UpdatedAt,
		})
	}
	return &reply.GetNotifies{List: replyNotifiesInfo, Total: int64(len(notifiesInfo))}, nil
}

func (notify) UpdateNotify(c *gin.Context, accountID uint, params *request.UpdateNotify) (*reply.UpdateNotify, errcode.Err) {
	qNotify := query.NewQueryNotify()
	IsLeader, err := qNotify.CheckIsLeader(accountID, params.RelationID)
	if err != nil {
		zap.S().Errorf("CreateNotify: qNotify.CheckIsLeader failes,err:%v", err)
		return nil, errcode.ErrServer
	}
	if !IsLeader {
		return nil, myerr.DoNotHaveAuth
	}
	notifyInfo, err := qNotify.UpdateNotify(params.ID, params.RelationID, params.MsgContent, params.MsgExpand)

	//TODO 通知所有人更新的 notify
	replyNotify := &reply.UpdateNotify{
		ID:         notifyInfo.ID,
		RelationID: notifyInfo.RelationID,
		MsgContent: notifyInfo.MsgContent,
		MsgExpand:  notifyInfo.MsgExpand,
		AccountID:  notifyInfo.AccountID,
		UpdateAt:   notifyInfo.UpdatedAt,
	}
	return replyNotify, nil
}
