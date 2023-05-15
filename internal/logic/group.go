/**
 * @Author: lenovo
 * @Description:
 * @File:  group
 * @Version: 1.0.0
 * @Date: 2023/05/07 17:49
 */

package logic

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-chat/internal/dao"
	"go-chat/internal/dao/mysql/query"
	tx2 "go-chat/internal/dao/mysql/tx"
	"go-chat/internal/global"
	"go-chat/internal/model/reply"
	"go-chat/internal/model/request"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app/errcode"
	"go-chat/internal/task"
	"go.uber.org/zap"
)

type vgroup struct{}

func (vgroup) CreateGroup(accountID uint, req request.CreateGroupReq) (*reply.CreateGroupReply, errcode.Err) {

	//生成setting表
	tx := tx2.NewGroupTX()
	rID, err := tx.CreateGroupSetting(accountID, req)
	if err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	rly := &reply.CreateGroupReply{
		Name:       req.Name,
		Avatar:     global.Settings.Rule.DefaultAccountAvatar,
		Signature:  req.SigNature,
		RelationID: rID,
		LeaderID:   accountID,
	}
	return rly, nil
}

func (vgroup) Dissolve(ctx *gin.Context, accountID uint, relationID int64) errcode.Err {

	//先要去判断一下解散群的是不是leader
	qS := query.NewSetting()
	settiingInfo, ok := qS.CheckRelationIDExist(int64(accountID), relationID)
	if !ok {
		return myerr.DoNotHaveThisRelation
	}
	//是否有权限
	if accountID != settiingInfo.AccountID {
		return myerr.DoNotHaveAuth
	}
	tx := tx2.NewGroupTX()
	if err := tx.Dissolve(relationID); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	tokenString := ctx.GetHeader(global.Settings.Token.AuthType)
	global.Worker.SendTask(task.DissolveGroup(tokenString, int64(settiingInfo.RelationID)))
	return nil
}

func (vgroup) Invite2Group(ctx *gin.Context, accountID int64, req request.InviteParamReq) errcode.Err {
	//先去判断一下accountID的relationID是否是存在的
	qGroup := query.NewGroup()
	ok := qGroup.ExistAccountIDAndRelationID(accountID, req.RelationID)
	if !ok {
		return errcode.ErrNotFound
	}
	//再去判断一下所选择的人是否是在群里面
	var ids []int64
	for _, id := range req.InvitePeopleIDs {
		if exist := qGroup.ExistAccountInGroup(req.RelationID, accountID); exist {
			continue
		}
		ids = append(ids, id)
	}
	tx := tx2.NewGroupTX()
	if err := tx.AddAccounts2GroupWithTX(accountID, req.RelationID, ids...); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	tokenString := ctx.GetHeader(global.Settings.Token.AuthType)
	global.Worker.SendTask(task.SendInvitedInfoToInviters(tokenString, ids))
	return nil
}

func (vgroup) GetGroupList(accountID int64) (*reply.GroupListReply, errcode.Err) {
	var rly reply.GroupListReply
	tx := tx2.NewGroupTX()
	GroupInfos, err := tx.GetGroupListTX(accountID)
	if err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	for _, groupInfo := range *GroupInfos {
		GroupItem := reply.GroupItem{
			Name:      groupInfo.Relation.GroupType.Name,
			Avatar:    groupInfo.Relation.GroupType.Avatar,
			Signature: groupInfo.Relation.GroupType.Signature,
			NickName:  groupInfo.NickName,
			IsDisturb: groupInfo.IsNotDisturbed,
			IsPin:     groupInfo.IsPin,
			IsShow:    groupInfo.IsShow,
		}
		rly.GroupItems = append(rly.GroupItems, GroupItem)
		rly.Total++
	}
	return &rly, nil
}

func (vgroup) TransferGroup(ctx *gin.Context, accountID int64, relationID int64, toID int64) errcode.Err {

	qS := query.NewSetting()
	_, ok := qS.CheckRelationIDExist(accountID, relationID)
	if !ok {
		return errcode.ErrServer
	}
	//先去判断一下是否是群主
	qGroup := query.NewGroup()
	if ok := qGroup.CheckIsLeader(accountID, relationID); !ok {
		return myerr.DoNotHaveAuth
	}
	//去查询一下，toID是否是在群里面
	if ok = qGroup.ExistAccountInGroup(relationID, toID); !ok {
		return myerr.DoNotHaveThisAccount
	}
	tx := tx2.NewGroupTX()
	if err := tx.TransferGroup(accountID, toID, relationID); err != nil {
		zap.S().Infof("TransferGroup failed,err:%v", err)
		return errcode.ErrServer.WithDetails(err.Error())
	}
	tokenString := ctx.GetHeader(global.Settings.Token.AuthType)

	global.Worker.SendTask(task.TransferGroup(tokenString, relationID))
	return nil
}

func (vgroup) QuitGroup(ctx *gin.Context, accountID int64, relationID int64) errcode.Err {
	//先去判断一下是否在群聊中
	ids, err := dao.Group.Redis.GetAllAccountsByRelationID(context.Background(), relationID)
	if err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	exist := false
	for _, id := range ids {
		if id == accountID {
			exist = true
			break
		}
	}
	if !exist {
		return myerr.DoNotHaveThisAccount
	}
	tx := tx2.NewGroupTX()
	if err := tx.QuitGroup(accountID, relationID); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	tokenString := ctx.GetHeader(global.Settings.Token.AuthType)
	global.Worker.SendTask(task.QuitGroup(tokenString, relationID))
	return nil
}

func (vgroup) GetMembers(relationID int64) (*reply.GetMembersReply, errcode.Err) {
	memberIds, err1 := dao.Group.Redis.GetAllAccountsByRelationID(context.Background(), relationID)
	if err1 != nil {
		zap.S().Infof("dao.Group.Redis.GetAllAccountsByRelationID failed: %v", err1)
		return nil, errcode.ErrServer.WithDetails(err1.Error())
	}
	qGroup := query.NewGroup()
	result, err := qGroup.GetSettingInfoByIDs(relationID, memberIds)
	if err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	return result, nil
}
