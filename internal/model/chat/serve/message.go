/**
 * @Author: lenovo
 * @Description:
 * @File:  message
 * @Version: 1.0.0
 * @Date: 2023/05/14 20:47
 */

package serve

import (
	"go-chat/internal/model/automigrate"
	"time"
)

type HandleSendMsg struct {
	AccountID   int64                  `json:"accountID"` //发送者ID
	AccessToken string                 `json:"accessToken"`
	RelationID  int64                  `json:"relation_id" ` // 关系ID
	MsgContent  string                 `json:"msg_content" ` // 消息内容
	MsgExtend   *automigrate.MsgExtend `json:"msg_extend"`   // 消息扩展信息
	RlyMsgID    int64                  `json:"rly_msg_id"`   // 回复消息ID (如果是回复消息，则此字段大于0)
}

type SendMsgRly struct {
	CreateTime time.Time `json:"create_time"` //发送消息的时间
	MsgID      int64     `json:"msg_id"`      //消息id
}
