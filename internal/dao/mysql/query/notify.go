package query

import (
	"fmt"
	"go-chat/internal/dao"
	"go-chat/internal/model/automigrate"
	"gorm.io/gorm"
)

type qNotify struct {
}

func NewQueryNotify() *qNotify {
	return &qNotify{}
}

func (qNotify) CheckIsLeader(accountID uint, relationID uint) (bool, error) {
	if result := dao.Group.DB.Model(&automigrate.Setting{}).Where(&automigrate.Setting{
		AccountID:  accountID,
		RelationID: relationID,
		IsLeader:   true,
	}).Find(&automigrate.Setting{}); result.Error != nil {
		return false, result.Error
	} else if result.RowsAffected == 0 {
		return false, fmt.Errorf("account is not Leader")
	}
	return true, nil
}

func (qNotify) CheckIsInGroup(accountID, relationID uint) (bool, error) {
	if result := dao.Group.DB.Model(&automigrate.Setting{}).Where(&automigrate.Setting{
		AccountID:  0,
		RelationID: 0,
	}).Find(&automigrate.Setting{}); result.Error != nil {
		return false, nil
	} else if result.RowsAffected == 0 {
		return false, fmt.Errorf("account is not in group")
	}
	return true, nil
}

func (qNotify) CreateNotify(accountID, relationID uint, msgContent string, msgExpand *automigrate.MsgExpand) (*automigrate.Notify, error) {
	notify := &automigrate.Notify{
		RelationID: relationID,
		AccountID:  accountID,
		MsgContent: msgContent,
		MsgExpand:  msgExpand,
	}
	if result := dao.Group.DB.Model(&automigrate.Notify{}).Create(notify); result.Error != nil {
		return nil, result.Error
	}
	return notify, nil
}

func (qNotify) DeleteNotify(notifyID, relationID uint) error {
	notify := &automigrate.Notify{
		Model:      gorm.Model{ID: notifyID},
		RelationID: relationID,
	}
	if result := dao.Group.DB.Model(&automigrate.Notify{}).Delete(notify); result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("deleteNotify failed, no notify deleted")
	}
	return nil
}

func (qNotify) UpdateNotify(notifyID, relationID uint, msgContent string, msgExpand *automigrate.MsgExpand) (*automigrate.Notify, error) {
	notifyInfo := &automigrate.Notify{}
	tx := dao.Group.DB.Begin()
	defer tx.Commit()
	updateNotify := &automigrate.Notify{
		MsgContent: msgContent,
		MsgExpand:  msgExpand,
	}
	//先更新数据
	if result := tx.Model(&automigrate.Notify{}).
		Where("id = ? AND relation_id = ?", notifyID, relationID).
		Updates(updateNotify); result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		tx.Rollback()
		return nil, fmt.Errorf("updateNotify failed, no notify updated")
	}
	//然后查询数据并返回
	if result := tx.Model(&automigrate.Notify{}).
		Where("id = ? AND relation_id = ?", notifyID, relationID).
		First(notifyInfo); result.Error != nil {
		tx.Rollback()
		return nil, fmt.Errorf("updateNotify failed, no notify updated")
	}

	return notifyInfo, nil
}

func (qNotify) GetNotifies(relationID uint) ([]automigrate.Notify, error) {
	notifies := make([]automigrate.Notify, 0)
	if result := dao.Group.DB.Model(&automigrate.Notify{}).
		Where("relation_id = ?", relationID).
		Find(&notifies); result.Error != nil {
		return nil, result.Error
	}
	return notifies, nil
}
