/**
 * @Author: lenovo
 * @Description:
 * @File:  message
 * @Version: 1.0.0
 * @Date: 2023/05/01 0:46
 */

package client

import (
	"go-chat/internal/model/automigrate"
)

type HandleSendMsgParams struct {
	RelationID int64                  `json:"relation_id" validate:"required,gte=1"`          // 关系ID
	MsgContent string                 `json:"msg_content" validate:"required,gte=1,lte=1000"` // 消息内容
	MsgExtend  *automigrate.MsgExtend `json:"msg_extend"`                                     // 消息扩展信息
	RlyMsgID   int64                  `json:"rly_msg_id"`                                     // 回复消息ID (如果是回复消息，则此字段大于0)
}
