/**
 * @Author: lenovo
 * @Description:
 * @File:  account
 * @Version: 1.0.0
 * @Date: 2023/03/27 22:30
 */

package automigrate

type Account struct {
	BaseModel
	UserID    uint
	User      User   `gorm:"foreignKey:UserID;references:ID"`
	Name      string `gorm:"type:varchar(255);not null"`
	Signature string `gorm:"type:varchar(255);not null"`
	Avatar    string `gorm:"type:varchar(255);not null" default:"http://lycmall.lyc666.xyz/chat/first.jpg"`
	Gender    string `gorm:"type:varchar(255);"`
}
