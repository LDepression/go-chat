package reply

import (
	"go-chat/internal/model/automigrate"
	"time"
)

type Notify struct {
	ID         uint                   `json:"id"`
	RelationID uint                   `json:"relation_id"`
	MsgContent string                 `json:"msg_content"`
	MsgExpand  *automigrate.MsgExpand `json:"msg_expand"`
	AccountID  uint                   `json:"account_id"`
	UpdateAt   time.Time              `json:"create_at"`
	//ReadIds    []uint                 `json:"read_ids"`
}

type GetNotifies struct {
	List  []Notify `json:"list"`
	Total int64    `json:"total,omitempty"`
}

type UpdateNotify struct {
	ID         uint                   `json:"id"`
	RelationID uint                   `json:"relation_id"`
	MsgContent string                 `json:"msg_content"`
	MsgExpand  *automigrate.MsgExpand `json:"msg_expand"`
	AccountID  uint                   `json:"account_id"`
	UpdateAt   time.Time              `json:"update_at"`
	//ReadIds    []uint                 `json:"read_ids"`
}
