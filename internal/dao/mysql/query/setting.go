package query

import (
	"fmt"
	"go-chat/internal/dao"
	"go-chat/internal/model/automigrate"
	"go-chat/internal/model/common"
)

type setting struct {
}

func NewQuerySetting() *setting {
	return &setting{}
}

func (setting) GetRelationInfoByAccountsID(selfAccountID, targetAccountID uint) (*automigrate.Relation, error) {
	var RelationInfo automigrate.Relation
	if result := dao.Group.DB.Model(&automigrate.Relation{}).
		Where("JSON_EXTRACT(friend_type,'$.AccountID1') = ? AND JSON_EXTRACT(friend_type,'$.AccountID2') = ?", selfAccountID, targetAccountID).
		Or("JSON_EXTRACT(friend_type,'$.AccountID1') = ? AND JSON_EXTRACT(friend_type,'$.AccountID2') = ?", targetAccountID, selfAccountID).
		First(&RelationInfo); result.Error != nil {
		return nil, result.Error
	}
	return &RelationInfo, nil
}

func (setting) GetRelationInfos(selfAccountID uint) ([]*automigrate.Relation, error) {
	var RelationInfos []*automigrate.Relation
	if result := dao.Group.DB.Model(&automigrate.Relation{}).
		Where("JSON_EXTRACT(friend_type,'$.AccountID1') = ? AND relation_type = ?", selfAccountID, common.RelationTypeFriend).
		Or("JSON_EXTRACT(friend_type,'$.AccountID2') = ? AND relation_type = ?", selfAccountID, common.RelationTypeFriend).
		Find(&RelationInfos); result.Error != nil {
		return nil, result.Error
	}
	return RelationInfos, nil
}

func (setting) GetRelationInfoByRelationID(relationID uint) (*automigrate.Relation, error) {
	var RelationInfo automigrate.Relation
	if result := dao.Group.DB.Model(&automigrate.Relation{}).
		Where("id = ?", relationID).
		First(&RelationInfo); result.Error != nil {
		return nil, result.Error
	}
	return &RelationInfo, nil
}

func (setting) GetFriendInfoByID(accountID, relationID uint) (*automigrate.Setting, error) {
	var FriendInfo automigrate.Setting
	if result := dao.Group.DB.Model(&automigrate.Setting{}).
		Where("relation_id = ?", relationID).
		Joins("Account").Where("account.id = ?", accountID).
		Joins("Relation").
		Find(&FriendInfo); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, nil
	}
	/*
		Where("relation_id = ?", relationID).
		Preload("Account", "id = ?", accountID).
		Preload("Relation", "id = ?", relationID).
		Find(&FriendInfo);
	*/
	return &FriendInfo, nil
}

func (setting) GetFriendInfoByName(accountID, relationID, limit, offset uint, name string) (*automigrate.Setting, error) {
	var FriendInfo automigrate.Setting
	if result := dao.Group.DB.Model(&automigrate.Setting{}).
		Where("settings.relation_id = ? AND settings.account_id = ? ", relationID, accountID).
		//TODO accounts.id 和 account.id不一样很烦
		Joins("Account").Where("account.id = ? AND (settings.nick_name LIKE ? OR account.name LIKE ?)", accountID, "%"+name+"%", "%"+name+"%").
		Preload("Relation").
		Preload("Account").
		Offset(int(offset)).Limit(int(limit)).
		Find(&FriendInfo); result.Error != nil {
		return nil, result.Error
		//Where("relation_id = ?", relationID).
		//Joins("left join settings on account.id = setting.account_id").
		//Joins("Account").Where("(account.name like ? OR setting.nick_name like ?) AND account.id = ?", "%"+name+"%", "%"+name+"%", accountID).
		//Joins("Relation").
		//Group("account.id"). // 如果 name既符合account.name和setting.nick_name，会产生重复数据
	} else if result.RowsAffected == 0 {
		return nil, nil
	}
	return &FriendInfo, nil
}

func (setting) UpdateNickName(targetAccountID, relationID uint, nickName string) error {
	if result := dao.Group.DB.Model(&automigrate.Setting{}).
		Where("relation_id = ? AND account_id = ?", relationID, targetAccountID).
		Update("nick_name", nickName); result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("UpdateNickName failed")
	}
	return nil
}

func (setting) UpdateSettingDisturb(targetAccountID, relationID uint, isDisturbed bool) error {
	var IsDisturbed int
	if isDisturbed {
		IsDisturbed = 1
	} else {
		IsDisturbed = 0
	}
	if result := dao.Group.DB.Model(&automigrate.Setting{}).
		Where("relation_id = ? AND account_id = ?", relationID, targetAccountID).
		Update("is_not_disturbed", IsDisturbed); result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("UpdateNickName failed")
	}
	return nil
}
