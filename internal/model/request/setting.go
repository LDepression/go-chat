package request

import "go-chat/internal/model/common"

type DeleteFriend struct {
	TargetAccountID uint `json:"target_account_id" form:"target_account_id"`
}

type GetFriendsByName struct {
	Name string `json:"name" form:"name"`
	common.Pager
}

type UpdateNickName struct {
	RelationID uint   `json:"relation_id" form:"relation_id"`
	NickName   string `json:"nick_name" form:"nick_name"`
}

type UpdateSettingDisturb struct {
	RelationID  uint `json:"relation_id" form:"relation_id"`
	IsDisturbed bool `json:"is_disturbed" form:"is_disturbed"`
}
