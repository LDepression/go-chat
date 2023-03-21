/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/03/20 21:17
 */

package query

import (
	"go-chat/internal/dao"
	"go-chat/internal/model/automigrate"
)

type quser struct {
}

func NewQueryUser() *quser {
	return &quser{}
}
func (quser) CheckEmailBeUsed(emailStr string) (bool, error) {
	var user automigrate.User
	result := dao.Group.DB.Where("email = ?", emailStr).Find(&user)
	return result.RowsAffected == 1, result.Error
}
