/**
 * @Author: lenovo
 * @Description:
 * @File:  setting
 * @Version: 1.0.0
 * @Date: 2023/04/04 18:06
 */

package query

import (
	"go-chat/internal/dao"
	"go-chat/internal/model/automigrate"
)

type Relation struct{}

func NewRelation() *Relation {
	return &Relation{}
}
func (Relation) CheckISFriend(account1ID uint64, account2ID uint64) bool {
	if result := dao.Group.DB.Model(&automigrate.Relation{}).Where("JSON_EXTRACT(friend_type, '$.AccountID1') = ? and JSON_EXTRACT(friend_type, '$.AccountID2') =?", account1ID, account2ID).Or(
		"JSON_EXTRACT(friend_type, '$.AccountID1') = ? and JSON_EXTRACT(friend_type, '$.AccountID2') =?", account2ID, account1ID,
	).Find(&Relation{}); result.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
