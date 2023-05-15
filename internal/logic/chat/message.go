/**
 * @Author: lenovo
 * @Description:
 * @File:  message
 * @Version: 1.0.0
 * @Date: 2023/05/14 20:39
 */

package chat

import (
	"go-chat/internal/dao/mysql/query"
	"go-chat/internal/global"
	"go-chat/internal/model/automigrate"
	"go-chat/internal/model/chat/serve"
	"go-chat/internal/model/reply"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app/errcode"
	"go-chat/internal/task"
	"go.uber.org/zap"
)

type message struct {
}

func (message) SendMsg(param serve.HandleSendMsg) (*serve.SendMsgRly, errcode.Err) {
	qS := query.NewSetting()
	_, ok := qS.CheckRelationIDExist(param.AccountID, param.RelationID)
	if !ok {
		return nil, myerr.DoNotHaveThisRelation
	}
	//如果是有回复的消息的话

	var rlyMsg *reply.RlyMsg
	if param.RlyMsgID > 0 {
		//先去判断一下RelationID是否是在群聊中
		qM := query.NewMessage()
		msgInfo, err := qM.CheckMsgInfoByID(param.RlyMsgID)
		if err != nil {
			zap.S().Infof("Error checking err：%v", err)
			return nil, errcode.ErrServer.WithDetails(err.Error())
		}
		if msgInfo.RelationID != param.RelationID {
			return nil, myerr.CanNotTalkInDifGroup
		}
		if msgInfo.IsRevoke {
			return nil, myerr.CanNotReplyRevoke
		}
		rlyMsg = &reply.RlyMsg{
			MsgID:      int64(msgInfo.ID),
			MsgExtend:  msgInfo.MsgExtend,
			MsgContent: msgInfo.MsgContent,
			IsRevoke:   msgInfo.IsRevoke,
			MsgType:    string(msgInfo.MsgType),
		}
	}

	global.Worker.SendTask(task.SendMsg(param.AccountID, reply.MsgInfoWithRly{
		MsgInfo: reply.MsgInfo{
			NotifyType: string(automigrate.MsgnotifytypeCommon),
			MsgType:    string(automigrate.MsgTypeText),
			MsgContent: param.MsgContent,
			Extend:     param.MsgExtend,
			AccountID:  param.AccountID,
			RelationID: param.RelationID,
		},
		RlyMsg: rlyMsg,
	}))
	return nil, nil
}
