/**
 * @Author: lenovo
 * @Description:
 * @File:  setting
 * @Version: 1.0.0
 * @Date: 2023/03/27 22:43
 */

package automigrate

import (
	"gorm.io/gorm"
	"time"
)

type Setting struct {
	gorm.Model
	AccountID      uint    // 目标id
	Account        Account `gorm:"foreignKey:AccountID;references:ID"`
	RelationID     uint
	Relation       Relation   `gorm:"foreignKey:RelationID;references:ID"`
	NickName       string     `gorm:"type:string;not null"`
	IsNotDisturbed bool       `gorm:"type:bool;not null"`
	IsPin          bool       `gorm:"type:bool;not null"`
	PinTime        *time.Time `gorm:"type:time"`
	IsShow         bool       `gorm:"type:bool;not null"`
	LastShowTime   *time.Time `gorm:"column:last_show_time;type:TIMESTAMP;default:CURRENT_TIMESTAMP  on update current_timestamp"`
	IsSelf         bool       `gorm:"type:bool;not null"`
	IsLeader       bool       `gorm:"type:bool;not null"`
}
