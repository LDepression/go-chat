package tx

import (
	"go-chat/internal/dao"
	"go-chat/internal/model/automigrate"
)

type settingTX struct {
}

func NewSettingTX() *settingTX {
	return &settingTX{}
}

func (settingTX) DeleteFriendWithTX(relationID uint) error {
	tx := dao.Group.DB.Begin()
	defer tx.Commit()
	// 通过relationID删除relation
	var relation automigrate.Relation
	if result := tx.Model(&automigrate.Relation{}).
		Where("id = ?", relationID).
		Delete(&relation); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// 通过relationID删除settings
	var setting automigrate.Setting
	if result := tx.Model(&automigrate.Setting{}).
		Where("relation_id = ?", relationID).
		Delete(&setting); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	return nil
}
