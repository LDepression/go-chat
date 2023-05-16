/**
 * @Author: lenovo
 * @Description:
 * @File:  message
 * @Version: 1.0.0
 * @Date: 2023/05/15 19:57
 */

package task

import (
	"context"
	"go-chat/internal/dao"
	"go-chat/internal/global"
	"go-chat/internal/model/chat"
	"go-chat/internal/model/reply"
	p2 "go-chat/internal/pkg/rocketmq/producer"
	"go.uber.org/zap"
)

func SendMsg(msg reply.MsgInfoWithRly) func() {
	return func() {
		ctx := context.Background()
		membs, err1 := dao.Group.Redis.GetAllAccountsByRelationID(ctx, msg.RelationID)
		if err1 != nil {
			zap.S().Infof("get all accounts failed err: %v", err1)
			return
		}
		for _, mID := range membs {
			if global.ChatMap.CheckIsOnConnection(mID) {
				global.ChatMap.Send(mID, chat.ClientSendMsg, msg)
			} else {
				p2.SendMsgToMQ(uint(mID))
			}
		}
		//}
	}

}
