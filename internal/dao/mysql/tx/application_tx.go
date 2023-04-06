package tx

import (
	"go-chat/internal/dao"
	"go-chat/internal/model/automigrate"
	"go-chat/internal/model/common"
	"gorm.io/gorm"
)

type ApplicationTX struct {
}

func NewApplicationTX() *ApplicationTX {
	return &ApplicationTX{}
}

func (ApplicationTX) AcceptApplicationWithTX(applicantID, receiverID uint) error {

	tx := dao.Group.DB.Begin()
	defer tx.Commit()
	// 将application中status更新为"已接受"
	if result := tx.Model(&automigrate.Application{}).
		Where("account_id1 = ? AND account_id2 = ?", applicantID, receiverID).
		Update("status", common.ApplicationStateAccepted); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// 建立两个人之间的relation
	var relation automigrate.Relation
	relation.RelationType = "friend"
	relation.FriendType = automigrate.FriendType{
		AccountID1: int64(applicantID),
		AccountID2: int64(receiverID),
	}
	if result := tx.Model(&automigrate.Relation{}).Create(&relation); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// 对两个account分别设置setting
	var settingAcceptInfo, settingReceiverInfo automigrate.Setting
	settingAcceptInfo.RelationID, settingReceiverInfo.RelationID = relation.ID, relation.ID
	//	增添对方id的setting
	settingAcceptInfo.AccountID, settingReceiverInfo.AccountID = receiverID, applicantID
	if result := tx.Model(&automigrate.Setting{}).Create(&settingAcceptInfo).Create(&settingReceiverInfo); result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return nil
}

func CommitTXFailed(tx *gorm.DB) error {
	if r := recover(); r != nil {
		tx.Rollback()
	}
	if err := tx.Commit().Error; err != nil {
		// 处理事务提交失败的错误
		return err
	}
	return nil
}
