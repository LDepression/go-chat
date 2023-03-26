/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/03/20 18:59
 */

package automigrate

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(200);not null;index:idx_email;unique"`
	Mobile   string `gorm:"type:varchar(100);not null;index:idx_mobile;unique"`
	Password string `gorm:"type:varchar(200);not null;"`
}
