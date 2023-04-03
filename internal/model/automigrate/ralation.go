/**
 * @Author: lenovo
 * @Description:
 * @File:  ralation
 * @Version: 1.0.0
 * @Date: 2023/03/27 21:18
 */

package automigrate

import "gorm.io/gorm"

type FriendType struct {
	AccountID1 int64
	AccountID2 int64
}

type GroupType struct {
	Name      string `gorm:"column:name"`
	Signature string `gorm:"column:signature"`
	Avatar    string `gorm:"avatar"`
}
type Relation struct {
	gorm.Model
	RelationType string     `gorm:"type:varchar(20);comment:关系类型(group/friend);not null"`
	FriendType   FriendType `gorm:"type:varchar(200);comment:好友类型存的值 例如:accountID1,accountID2;not null;"`
	GroupType    GroupType  `gorm:"type:varchar(200);comment:群组的类型存的值:例如:群组名字,群签名,群头像;not null;"`
}
