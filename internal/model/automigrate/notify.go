package automigrate

import "gorm.io/gorm"

type Notify struct {
	gorm.Model
	RelationID uint       `json:"relation_id"`
	AccountID  uint       `json:"account_id"`
	MsgContent string     `json:"msg_content" gorm:"type:varchar(800)"`
	MsgExpand  *MsgExpand `gorm:"embedded" json:"msg_expand"`
}
