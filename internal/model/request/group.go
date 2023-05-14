/**
 * @Author: lenovo
 * @Description:
 * @File:  group
 * @Version: 1.0.0
 * @Date: 2023/05/07 17:46
 */

package request

type CreateGroupReq struct {
	SigNature string `json:"SigNature" binding:"required"`
	Name      string `json:"Name" binding:"required"`
	Avatar    string `json:"Avatar" `
}

type DissolveReq struct {
	RelationID int64 `json:"RelationID" binding:"required,omitempty"`
}

type InviteParamReq struct {
	InvitePeopleIDs []int64 `json:"InvitePeopleIDs" binding:"required"`
	RelationID      int64   `json:"RelationID" binding:"required"`
}

type TransferReq struct {
	ToID       int64 `json:"ToID" binding:"required"`
	RelationID int64 `json:"RelationID" binding:"required"`
}

type BeQuitedReq struct {
	RelationID int64 `json:"RelationID" binding:"required"`
	AccountID  int64 `json:"AccountID" binding:"required"`
}

type GetMems struct {
	RelationID int64 `json:"RelationID" binding:"required"`
}

type QuitReq struct {
	RelationID int64 `json:"RelationID" binding:"required"`
}
