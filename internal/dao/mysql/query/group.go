/**
 * @Author: lenovo
 * @Description:
 * @File:  group
 * @Version: 1.0.0
 * @Date: 2023/05/07 17:56
 */

package query

import (
	"errors"
	"go-chat/internal/dao"
	"go-chat/internal/global"
	"go-chat/internal/model/automigrate"
	"go-chat/internal/model/reply"
	"go-chat/internal/model/request"
	"gorm.io/gorm"
	"sync"
)

type group struct{}

func NewGroup() *group {
	return &group{}
}

var (
	ErrorCreateData = errors.New("插入数据失败")
)

func (group) CreateGroupRelation(req request.CreateGroupReq) (uint, error) {
	relation := &automigrate.Relation{
		RelationType: "group",
		GroupType: automigrate.GroupType{
			Name:      req.Name,
			Signature: req.SigNature,
			Avatar:    global.Settings.Rule.DefaultAccountAvatar,
		},
	}
	if result := dao.Group.DB.Model(&automigrate.Relation{}).Create(relation); result.RowsAffected == 0 {
		return 0, ErrorCreateData
	}
	return relation.ID, nil
}

func (group) ExistAccountIDAndRelationID(accountID int64, relationID int64) bool {
	settingInfo := automigrate.Setting{}
	if result := dao.Group.DB.Model(&automigrate.Setting{}).Where("account_id =? AND relation_id = ?", accountID, relationID).Find(&settingInfo); result.RowsAffected == 0 {
		return false
	}
	return true
}

func (group) ExistAccountInGroup(relationID int64, accountID int64) bool {
	return dao.Group.DB.Model(&automigrate.Setting{}).Where("account_id = ?  AND relation_id =?", accountID, relationID).Find(&automigrate.Setting{}).RowsAffected != 0
}

func (group) CheckIsLeader(accountID int64, relationID int64) bool {
	return dao.Group.DB.Model(&automigrate.Setting{}).Where(&automigrate.Setting{
		AccountID:  uint(accountID),
		RelationID: uint(relationID),
		IsLeader:   true,
	}).Find(&automigrate.Setting{}).RowsAffected != 0
}

func (group) GetSettingInfoByIDs(relationID int64, memberIDs []int64) (*reply.GetMembersReply, error) {

	type Result struct {
		ID        int64  `json:"id"`
		Name      string `json:"name"`
		Avatar    string `json:"avatar"`
		Signature string `json:"signature"`
		Total     int64  `json:"total"`
	}
	var results []Result
	dao.Group.DB.Raw(`
SELECT
	a.id,
	a.avatar,
	a.name,
	a.signature,
	(
	SELECT
		COUNT(*) 
	FROM
		settings s 
	WHERE
		s.relation_id = ? 
	AND s.account_id IN ? 
	) AS total 
FROM
	accounts a 
WHERE
	a.id IN  ? ;
`, relationID, memberIDs, memberIDs).Scan(&results)
	if len(results) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	var rly reply.GetMembersReply
	once := sync.Once{}
	for _, m := range results {
		rly.MembersInfo = append(rly.MembersInfo, reply.MemberInfo{
			AccountID: m.ID,
			Name:      m.Name,
			Avatar:    m.Avatar,
		})
		once.Do(func() {
			rly.Total = m.Total
		})
	}
	return &rly, nil
}
