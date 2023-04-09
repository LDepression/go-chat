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
	"go-chat/internal/model/common"
)

type qApplication struct {
}

func NewQueryApplication() *qApplication {
	return &qApplication{}
}

func (qApplication) GetApplicationByID(applicantID, receiverID uint) (*automigrate.Application, error) {
	var ApplicationInfo automigrate.Application
	if result := dao.Group.DB.Model(&automigrate.Application{}).
		Where("account_id1 = ? AND account_id2 = ?", applicantID, receiverID).
		First(&ApplicationInfo); result.Error != nil {
		return nil, result.Error
	}
	return &ApplicationInfo, nil
}

func (qApplication) RefuseApplication(applicantID, receiverID uint, refuseMsg string) error {
	values := map[string]interface{}{
		"status":     common.ApplicationStateRefused,
		"refuse_msg": refuseMsg,
	}

	if result := dao.Group.DB.Model(&automigrate.Application{}).
		Where("account_id1 = ? AND account_id2 = ?", applicantID, receiverID).
		Updates(values); result.Error != nil {
		return result.Error
	}
	return nil
}

func (qApplication) GetApplicationsList(accountID, limit, offset uint) ([]*automigrate.Application, int64, error) {
	var applicationList []*automigrate.Application
	var totalCount int64
	if result := dao.Group.DB.Model(&automigrate.Application{}).
		Where("account_id1 = ? OR account_id2 = ?", accountID, accountID).
		Offset(int(offset)).Limit(int(limit)).
		Find(&applicationList); result.Error != nil {
		return nil, 0, result.Error
	} else {
		totalCount = result.RowsAffected
	}
	return applicationList, totalCount, nil
}

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
