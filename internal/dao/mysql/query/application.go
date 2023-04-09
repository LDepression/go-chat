/**
 * @Author: lenovo
 * @Description:
 * @File:  application
 * @Version: 1.0.0
 * @Date: 2023/04/06 19:41
 */

package query

import (
	"go-chat/internal/dao"
	"go-chat/internal/model/automigrate"
	"go-chat/internal/model/request"
)

type application struct {
}

func NewApplication() *application {
	return &application{}
}
func (application) CreateApplication(ApplicationID1 uint64, req request.CreateApplicationReq) (uint64, error) {
	var application automigrate.Application
	application.AccountID1 = ApplicationID1
	application.AccountID2 = req.AccountID
	application.ApplyMsg = req.ApplicationMsg
	application.Status = string(WAITING)
	result := dao.Group.DB.Model(&automigrate.Application{}).Create(&application)
	return application.ID, result.Error
}

func (application) DeleteApplication(accountID uint64) error {
	result := dao.Group.DB.Model(&automigrate.Application{}).Where(automigrate.Application{AccountID2: accountID}).Delete(&automigrate.Application{})
	return result.Error
}
