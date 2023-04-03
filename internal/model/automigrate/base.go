/**
 * @Author: lenovo
 * @Description:
 * @File:  base
 * @Version: 1.0.0
 * @Date: 2023/03/29 20:34
 */

package automigrate

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint64         `gorm:"type:bigint"`
	CreatedAt time.Time      `gorm:"column:add_time"`
	UpdatedAt time.Time      `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt `gorm:"column:delete_time"`
	IsDelete  *bool          `gorm:"column:is_deleted"`
}
