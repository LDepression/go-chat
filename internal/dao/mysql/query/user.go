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

func (quser) SaveRegisterInfo(user automigrate.User) (uint, error) {
	result := dao.Group.DB.Create(&user)
	return user.ID, result.Error
}

func (quser) GetUserByEmail(emailStr string) (*automigrate.User, error) {
	var user automigrate.User
	if result := dao.Group.DB.Where(automigrate.User{Email: emailStr}).Find(&user); result.RowsAffected == 0 {
		return nil, result.Error
	}
	return &user, nil
}

func (quser) GetUserByID(userID int64) (*automigrate.User, error) {
	var user automigrate.User
	if result := dao.Group.DB.Model(&automigrate.User{}).Where("id =?", userID).Find(&user); result.RowsAffected == 0 {
		return nil, result.Error
	}
	return &user, nil
}

func (quser) ModifyPassword(email string, hashPassword string) error {
	result := dao.Group.DB.Model(&automigrate.User{}).Where(automigrate.User{Email: email}).Update("password", hashPassword)
	return result.Error
}
