/**
 * @Author: lenovo
 * @Description:
 * @File:  setting
 * @Version: 1.0.0
 * @Date: 2023/04/18 21:00
 */

package reply

import "go-chat/internal/model/automigrate"

type FriendsInfo struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type FriendType struct {
	AccountID1 int64
	AccountID2 int64
}

type GroupInfo struct {
	Name      string `gorm:"column:name" json:"Name"`
	Signature string `gorm:"column:signature" json:"Signature"`
	Avatar    string `gorm:"avatar" json:"Avatar"`
}

type SettingInfo struct {
	AccountID      int64  `json:"account_id"`
	RelationID     int64  `json:"relation_id"`
	NickName       string `json:"nick_name"`
	IsNotDisturbed bool   `json:"is_not_disturbed"`
	IsPin          bool   `json:"is_pin"`
	IsShow         bool   `json:"is_show"`
}
type GetSettingRep struct {
	Friend      *FriendsInfo
	Group       *GroupInfo
	BaseSetting automigrate.Setting
}

type SettingReq struct {
	Data  []GetSettingRep
	Total int64
}
