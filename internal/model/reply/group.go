/**
 * @Author: lenovo
 * @Description:
 * @File:  group
 * @Version: 1.0.0
 * @Date: 2023/05/07 19:37
 */

package reply

type CreateGroupReply struct {
	Name       string `json:"Name" `
	Avatar     string `json:"Avatar" `
	Signature  string `json:"Signature" `
	RelationID uint   `json:"RelationID" `
	LeaderID   uint   `json:"LeaderID" `
}

type GroupItem struct {
	Name      string `json:"Name"`
	Avatar    string `json:"Avatar"`
	Signature string `json:"Signature"`
	NickName  string `json:"NickName"`
	IsDisturb bool   `json:"IsDisturb"`
	IsPin     bool   `json:"IsPin"`
	IsShow    bool   `json:"IsShow"`
}

type GroupListReply struct {
	Total      int64       `json:"Total"`
	GroupItems []GroupItem `json:"GroupItems"`
}

type MemberInfo struct {
	AccountID int64  `json:"AccountID"`
	Name      string `json:"Name"`
	Avatar    string `json:"Avatar"`
}

type GetMembersReply struct {
	Total       int64        `json:"Total"`
	MembersInfo []MemberInfo `json:"MembersInfo"`
}
