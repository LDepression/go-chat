/**
 * @Author: lenovo
 * @Description:
 * @File:  message
 * @Version: 1.0.0
 * @Date: 2023/05/14 23:11
 */

package reply

import (
	"go-chat/internal/model/automigrate"
	"time"
)

// RlyMsg 回复消息详情 可能为null
type RlyMsg struct {
	MsgID      int64                  `json:"msg_id"`      // 回复消息ID
	MsgType    string                 `json:"msg_type"`    // 消息类型 [text,file]
	MsgContent string                 `json:"msg_content"` // 消息内容 文件则为url，文本则为文本内容，由拓展信息进行补充
	MsgExtend  *automigrate.MsgExtend `json:"msg_extend"`  // 消息扩展信息 可能为null
	IsRevoke   bool                   `json:"is_revoke"`   // 是否撤回
}

// MsgInfo 完整的消息详情
type MsgInfo struct {
	ID         int64                  `json:"id"`          // 消息ID
	NotifyType string                 `json:"notify_type"` // 通知类型 [system,common]
	MsgType    string                 `json:"msg_type"`    // 消息类型 [text,file]
	MsgContent string                 `json:"msg_content"` // 消息内容 文件则为url，文本则为文本内容，由拓展信息进行补充
	Extend     *automigrate.MsgExtend `json:"msg_extend"`  // 消息扩展信息 可能为null
	FileID     int64                  `json:"file_id"`     // 文件ID 当消息类型为file时>0
	AccountID  int64                  `json:"account_id"`  // 账号ID 发送者ID
	RelationID int64                  `json:"relation_id"` // 关系ID
	CreateAt   time.Time              `json:"create_at"`   // 创建时间
	IsRevoke   bool                   `json:"is_revoke"`   // 是否撤回
	IsTop      bool                   `json:"is_top"`      // 是否置顶
	IsPin      bool                   `json:"is_pin"`      // 是否pin
	PinTime    time.Time              `json:"pin_time"`    // pin时间
	ReadIds    []int64                `json:"read_ids"`    // 已读的账号ID 当请求者不为发送者时为空
	ReplyCount int64                  `json:"reply_count"` // 回复数
}

// MsgInfoWithRly 完整的消息详情 包含回复消息
type MsgInfoWithRly struct {
	MsgInfo
	RlyMsg *RlyMsg `json:"rly_msg"` // 回复消息详情 可能为null
}
