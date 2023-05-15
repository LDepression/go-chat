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
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"go-chat/internal/dao"
	"go-chat/internal/dao/mysql/query"
	"go-chat/internal/global"
	"go-chat/internal/model/automigrate"
	"go-chat/internal/model/chat"
	"go-chat/internal/model/reply"
	"go.uber.org/zap"
)

func SendMsg(accountID int64, msg reply.MsgInfoWithRly) func() {
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
				p, err := rocketmq.NewProducer(producer.WithNameServer([]string{"192.168.28.30:9876"}))
				if err != nil {
					panic("生成producer失败")
				}
				if err = p.Start(); err != nil {
					panic(err)
				}
				uID := fmt.Sprintf("accountID:%d", mID)
				MsgJson, err := json.Marshal(msg)
				if err != nil {
					zap.S().Infof("marshal error err: %v", err)
					return
				}
				res, err := p.SendSync(context.Background(), primitive.NewMessage(uID, MsgJson))
				if err != nil {
					fmt.Println("发送失败", err)
				} else {
					fmt.Println("发送成功,res:", res.String())
				}
				if err := p.Shutdown(); err != nil {
					panic(err)
				}
			}
			// todo:此时将数据入库
		}
		var rlyMsgID int64
		if msg.RlyMsg != nil {
			rlyMsgID = msg.RlyMsg.MsgID
		}
		param := query.CreateMsgParams{
			AccountID:  sql.NullInt64{Int64: accountID, Valid: true},
			MsgType:    string(automigrate.MsgTypeText),
			RelationID: sql.NullInt64{Int64: msg.RelationID, Valid: true},
			MsgContent: msg.MsgContent,
			MsgExtend:  msg.Extend,
			ReplyMsgID: sql.NullInt64{Int64: rlyMsgID, Valid: rlyMsgID > 0},
		}
		qM := query.NewMessage()
		msgID, err := qM.CreateMsg(param)
		if err != nil {
			zap.S().Infof("create message %v", err)
			return
		} else {
			fmt.Printf("MsgID------------------------------> %d", msgID)
		}
		//}
	}

}
