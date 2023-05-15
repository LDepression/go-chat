/**
 * @Author: lenovo
 * @Description:
 * @File:  message
 * @Version: 1.0.0
 * @Date: 2023/05/14 20:57
 */

package automigrate

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Remind struct {
	Idx       int64 `json:"idx" binding:"required,gte=1" validate:"required,gte=1"`        // 第几个@
	AccountID int64 `json:"account_id" binding:"required,gte=1" validate:"required,gte=1"` // 被@的账号ID
}

// MsgExtend 消息扩展信息 可能为null
type MsgExtend struct {
	Remind []Remind `json:"remind"` // @的描述信息
}

type Msgnotifytype string

const (
	MsgnotifytypeSystem Msgnotifytype = "system"
	MsgnotifytypeCommon Msgnotifytype = "common"
)

type MsgType string

const (
	MsgTypeText MsgType = "text"
	MsgTypeFile MsgType = "file"
)

type Message struct {
	gorm.Model
	NotifyType Msgnotifytype `gorm:"type:varchar(20);not null"`
	MsgType    MsgType       `gorm:"type:varchar(200);not null"`
	MsgContent string        `json:"msg_content"`
	MsgExtend  *MsgExtend    `json:"msg_extend"`
	FileID     sql.NullInt64
	File       File `gorm:"foreignKey:FileID;references:ID"`
	AccountID  sql.NullInt64
	Account    Account `gorm:"foreignKey:AccountID;references:ID"`
	RlyMsgID   sql.NullInt64
	Message    *Message `gorm:"foreignKey:RlyMsgID;references:ID"`
	RelationID int64
	Relation   Relation   `gorm:"foreignKey:RelationID;references:ID"`
	IsRevoke   bool       `gorm:"type:boolean"`
	IsTop      bool       `gorm:"type:boolean"`
	IsPin      bool       `gorm:"type:boolean"`
	PinTime    *time.Time `gorm:"type:time"`
	ReadIds    []int64    `gorm:"type:text"`
}
