package automigrate

// MsgExpand 消息扩展信息 可能为null
type MsgExpand struct {
	Reminds []Remind `gorm:"column:reminds;type:varchar(255);serializer:json" json:"remind"` // @的描述信息
}

type Remind struct {
	Idx       uint `json:"idx" binding:"required,gte=1" validate:"required,gte=1"`        // 第几个@
	AccountID uint `json:"account_id" binding:"required,gte=1" validate:"required,gte=1"` // 被@的账号ID
}
