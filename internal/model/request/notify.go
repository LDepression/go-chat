package request

import "go-chat/internal/model/automigrate"

type CreateNotify struct {
	RelationID uint                   `json:"relation_id" form:"relation_id" binding:"required"`
	MsgContent string                 `json:"msg_content" form:"msg_content" binding:"required"`
	MsgExpand  *automigrate.MsgExpand `json:"msg_expand"  form:"msg_expand" `
}

type DeleteNotify struct {
	ID         uint `json:"notify_id" form:"notify_id" binding:"required"`
	RelationID uint `json:"relation_id" form:"relation_id" binding:"required"`
}

type GetNotifyByID struct {
	RelationID uint `json:"relation_id" form:"relation_id" binding:"required"`
}

type UpdateNotify struct {
	ID         uint                   `json:"notify_id" form:"notify_id" binding:"required"`
	RelationID uint                   `json:"relation_id" form:"relation_id" binding:"required"`
	MsgContent string                 `json:"msg_content" form:"msg_content" binding:"required"`
	MsgExpand  *automigrate.MsgExpand ` json:"msg_expand"  form:"msg_expand"`
}
