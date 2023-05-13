/**
 * @Author: lenovo
 * @Description:
 * @File:  setting_tx
 * @Version: 1.0.0
 * @Date: 2023/04/18 21:11
 */

package tx

import (
	"errors"
	"go-chat/internal/dao"
	"go-chat/internal/model/automigrate"
	"go-chat/internal/model/reply"
	"gorm.io/gorm"
)

type SettingTX struct{}

func NewSettingTX() *SettingTX {
	return &SettingTX{}
}

func (SettingTX) GetFriendsPinsInfo(accountID uint64) (*reply.SettingReq, error) {
	tx := dao.Group.DB.Begin()
	//先要去relation表中查询出来friends
	var relationIDs []uint
	if result := tx.Raw(`
	SELECT
	r.id 
FROM
	relations r,
	settings s
WHERE
	(
		JSON_VALID( friend_type ) 
	AND ( JSON_EXTRACT( friend_type, '$.AccountID1' ) = ? OR JSON_EXTRACT( friend_type, '$.AccountID2' ) = ? )) 
	AND r.deleted_at IS NULL
	AND s.relation_id = r.id
	
ORDER BY
	s.pin_time DESC
`, accountID, accountID).Scan(&relationIDs); result.RowsAffected == 0 {
		return nil, errors.New("该账号暂无好友")
	}

	var SettingRows []automigrate.Setting
	if result := tx.Model(&automigrate.Setting{}).Where("relation_id IN (?) AND is_pin = ?", relationIDs, 1).Preload("Account").Find(&SettingRows); result.RowsAffected == 0 {
		return nil, result.Error
	}
	tx.Commit()
	replySettingResult := reply.SettingReq{}
	//接下来去account表中去查询相关信息
	for _, friendInfo := range SettingRows {
		FriendsPinsInfo := &reply.GetSettingRep{}
		FriendsPinsInfo.BaseSetting = friendInfo
		FriendsPinsInfo.Group = nil
		FInfo := &reply.FriendsInfo{}
		FInfo.Avatar = friendInfo.Account.Avatar
		FInfo.Gender = friendInfo.Account.Gender
		FInfo.Name = friendInfo.Account.Name
		FriendsPinsInfo.Friend = FInfo
		replySettingResult.Data = append(replySettingResult.Data, *FriendsPinsInfo)
		replySettingResult.Total++
	}
	return &replySettingResult, nil
}

func (SettingTX) GetGroupsPinsInfo(accountID uint64) (*reply.SettingReq, error) {
	tx := dao.Group.DB.Begin()
	//先去把这个人加的群给找出来
	var relationIDs []int64
	tx.Preload("Account").
		Preload("Relation").
		Raw(`
SELECT
	settings.relation_id
FROM
	settings
	LEFT JOIN relations ON settings.relation_id = relations.id 
WHERE
	settings.account_id = ? 
	AND relations.relation_type = "group";
`, accountID).Scan(&relationIDs)
	//先去找到setting相关的信息
	var GroupSetting []automigrate.Setting
	result := tx.Model(&automigrate.Setting{}).Where("relation_id IN(?) AND account_id = ?", relationIDs, accountID).Order("pin_time desc").Preload("Relation").Find(&GroupSetting)
	if result.Error != nil {
		return nil, result.Error
	}
	tx.Commit()
	//再去查找群组的相关信息
	var resultReply reply.SettingReq
	for _, GroupInfo := range GroupSetting {
		var replySettingReply reply.GetSettingRep
		replySettingReply.BaseSetting = GroupInfo
		replySettingReply.Friend = nil
		GInfo := &reply.GroupInfo{}
		GInfo.Avatar = GroupInfo.Relation.GroupType.Avatar
		GInfo.Name = GroupInfo.Relation.GroupType.Name
		GInfo.Signature = GroupInfo.Relation.GroupType.Signature
		replySettingReply.Group = GInfo
		resultReply.Data = append(resultReply.Data, replySettingReply)
	}
	resultReply.Total = int64(len(resultReply.Data))

	return &resultReply, nil
}

func (SettingTX) GetFriendsShowsOrderByShowTime(accountID uint64) (*reply.SettingReq, error) {
	tx := dao.Group.DB.Begin()

	var rids []int64
	if err := tx.Raw(`
SELECT
	DISTINCT r.id
FROM
	relations r
	JOIN settings s ON r.id = s.relation_id
WHERE
	JSON_VALID( r.friend_type ) 
	AND ( JSON_EXTRACT( r.friend_type, '$.AccountID1' ) = ? OR JSON_EXTRACT( r.friend_type, '$.AccountID2' ) = ? ) 
	AND relation_type = "friend"
	AND r.deleted_at IS NULL
`, accountID, accountID).Find(&rids); err.Error != nil {
		return nil, err.Error
	}
	var resultInfo []automigrate.Setting
	if result := tx.Model(&automigrate.Setting{}).
		Joins("JOIN relations ON settings.relation_id = relations.id").
		Joins("JOIN accounts ON settings.account_id = accounts.id ").
		Where("settings.relation_id IN(?) and is_show =? and account_id <>?", rids, 1, accountID).Preload("Relation").Preload("Account").Find(&resultInfo); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	tx.Commit()
	//接下来就直接从

	replySettingResult := reply.SettingReq{}
	//接下来去account表中去查询相关信息
	for _, sInfo := range resultInfo {
		settingInfo := sInfo
		accountInfo := sInfo.Account
		FriendsPinsInfo := &reply.GetSettingRep{}
		FriendsPinsInfo.BaseSetting = settingInfo
		FriendsPinsInfo.Group = nil
		FInfo := &reply.FriendsInfo{}
		FInfo.Avatar = accountInfo.Avatar
		FInfo.Gender = accountInfo.Gender
		FInfo.Name = accountInfo.Name
		FriendsPinsInfo.Friend = FInfo
		replySettingResult.Data = append(replySettingResult.Data, *FriendsPinsInfo)
		replySettingResult.Total++
	}
	return &replySettingResult, nil
}

func (SettingTX) GetGroupShowsOrderBy(accountID uint64) (*reply.SettingReq, error) {
	//先去查询一下对应的群组
	tx := dao.Group.DB.Begin()
	var settingInfos []automigrate.Setting
	var rIds []int64
	if err := tx.Raw(`
	SELECT
		s.relation_id 
	FROM
		settings s
		LEFT JOIN relations r ON s.relation_id = r.id 
	WHERE
		( s.account_id = ? AND r.relation_type = "group" )
`, accountID).Scan(&rIds); err.Error != nil {
		return nil, err.Error
	}
	if result := tx.Model(&automigrate.Setting{}).Joins("JOIN relations ON settings.relation_id = relations.id ").
		Where("settings.relation_id IN(?) AND settings.account_id =?", rIds, accountID).Preload("Relation").Order("settings.last_show_time DESC").
		Find(&settingInfos); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	tx.Commit()
	//再去查找群组的相关信息
	var resultReply reply.SettingReq
	for _, GroupInfo := range settingInfos {
		SettingInfo := GroupInfo
		RelationInfo := GroupInfo.Relation
		var replySettingReply reply.GetSettingRep
		replySettingReply.BaseSetting = SettingInfo
		replySettingReply.Friend = nil
		GInfo := &reply.GroupInfo{}
		GInfo.Avatar = RelationInfo.GroupType.Avatar
		GInfo.Name = RelationInfo.GroupType.Name
		GInfo.Signature = GroupInfo.Relation.GroupType.Signature
		replySettingReply.Group = GInfo
		resultReply.Data = append(resultReply.Data, replySettingReply)
	}
	resultReply.Total = int64(len(resultReply.Data))

	return &resultReply, nil
}
