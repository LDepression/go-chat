/**
 * @Author: lenovo
 * @Description:
 * @File:  setting
 * @Version: 1.0.0
 * @Date: 2023/05/05 17:11
 */

package query

import (
	"errors"
	"go-chat/internal/dao"
	"go-chat/internal/model/automigrate"
)

type setting struct{}

func NewSetting() *setting {
	return &setting{}
}

var ErrorNoUpdateRow = errors.New("暂无更新的部分")

func (setting) UpdatePinsState(relationID int64, isPin bool) error {
	if result := dao.Group.DB.Model(&automigrate.Setting{}).Where("relation_id = ?", relationID).Update("is_pin", isPin); result.RowsAffected == 0 {
		return ErrorNoUpdateRow
	}
	return nil
}

func (setting) CheckRelationIDExist(accountID, relationID int64) (automigrate.Setting, bool) {
	var settingInfo automigrate.Setting
	if result := dao.Group.DB.Model(&automigrate.Setting{}).Where("relation_id = ? AND account_id = ?", relationID, accountID).Find(&settingInfo); result.RowsAffected == 0 {
		return settingInfo, false
	}
	return settingInfo, true
}

func (setting) UpdateNickName(NickName string, relationID int64) error {
	if result := dao.Group.DB.Model(&automigrate.Setting{}).Where("relation_id = ?", relationID).Update("nick_name", NickName); result.RowsAffected == 0 {
		return ErrorNoUpdateRow
	}
	return nil
}

func (setting) UpdateIsDisturbState(relationID int64, isDisturbState bool) error {
	if result := dao.Group.DB.Model(&automigrate.Setting{}).Where("relation_id = ?", relationID).Update("is_not_disturbed", isDisturbState); result.RowsAffected == 0 {
		return ErrorNoUpdateRow
	}
	return nil
}

func (setting) UpdateIsShowState(relatiionID int64, isShowState bool) error {
	if result := dao.Group.DB.Model(&automigrate.Setting{}).Where("relation_id = ?", relatiionID).Update("is_pin", isShowState); result.RowsAffected == 0 {
		return ErrorNoUpdateRow
	}
	return nil
}
