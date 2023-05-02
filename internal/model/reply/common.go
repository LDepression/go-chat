/**
 * @Author: lenovo
 * @Description:
 * @File:  common
 * @Version: 1.0.0
 * @Date: 2023/03/28 22:53
 */

package reply

import (
	"go-chat/internal/model/common"
	"time"
)

type EasyAccount struct {
	AccountID uint   `json:"account_id" gorm:"type:bigint"`
	Name      string `json:"name" gorm:"type:varchar(255);not null"`
	Avatar    string `json:"avatar" gorm:"type:varchar(255);not null"`
}

type EasySetting struct {
	RelationID     uint                `gorm:"relation_id" json:"relation_id"`
	RelationType   common.RelationType `gorm:"type:varchar(20);not null" json:"relation_type"`
	NickName       string              `gorm:"type:string;not null" json:"nick_name"`
	IsNotDisturbed bool                `gorm:"type:bool;not null" json:"is_not_disturbed"`
	IsPin          bool                `gorm:"type:bool;not null" json:"is_pin"`
	PinTime        time.Time           `gorm:"type:time;not null" json:"pin_time"`
	IsShow         bool                `gorm:"type:bool;not null" json:"is_show"`
	LastShowTime   time.Time           `gorm:"column:last_show_time;type:TIMESTAMP;default:CURRENT_TIMESTAMP  on update current_timestamp" json:"last_show_time"`
	IsSelf         bool                `gorm:"type:bool;not null"`
}
