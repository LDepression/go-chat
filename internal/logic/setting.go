/**
 * @Author: lenovo
 * @Description:
 * @File:  setting
 * @Version: 1.0.0
 * @Date: 2023/04/18 20:59
 */

package logic

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/dao/mysql/query"
	tx2 "go-chat/internal/dao/mysql/tx"
	"go-chat/internal/global"
	"go-chat/internal/model/chat/serve"
	"go-chat/internal/model/reply"
	"go-chat/internal/model/request"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app/errcode"
	"go-chat/internal/task"
	"go.uber.org/zap"
)

type setting struct{}

func (setting) GetPins(accountID uint64) (*reply.SettingReq, errcode.Err) {
	tx := tx2.NewSettingTX()
	frindsPinsRows, err := tx.GetFriendsPinsInfo(accountID)
	if err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	groupsPinsRows, err := tx.GetGroupsPinsInfo(accountID)
	if err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}

	//将上面的按照pin的时间来进行一个排序
	var totalSettingReq reply.SettingReq

	//两个pinTime都是按照降序排列的，所以直接使用双指针进行排序
	var i, j int
	for i, j = 0, 0; i < int(frindsPinsRows.Total) && j < int(groupsPinsRows.Total); {
		if frindsPinsRows.Data[i].BaseSetting.PinTime.After(*groupsPinsRows.Data[j].BaseSetting.PinTime) {
			totalSettingReq.Data = append(totalSettingReq.Data, frindsPinsRows.Data[i])
			totalSettingReq.Total++
			i++
		} else {
			totalSettingReq.Data = append(totalSettingReq.Data, groupsPinsRows.Data[j])
			totalSettingReq.Total++
			j++
		}
	}
	for i < int(frindsPinsRows.Total) {
		totalSettingReq.Data = append(totalSettingReq.Data, frindsPinsRows.Data[i])
		totalSettingReq.Total++
		i++
	}
	for j < int(groupsPinsRows.Total) {
		totalSettingReq.Data = append(totalSettingReq.Data, groupsPinsRows.Data[j])
		totalSettingReq.Total++
		j++
	}

	totalSettingReq.Total = frindsPinsRows.Total + groupsPinsRows.Total

	return &totalSettingReq, nil
}

func (setting) GetShowsOrderByShowTime(accountID uint64) (*reply.SettingReq, errcode.Err) {
	//展示在首页上面
	tx := tx2.NewSettingTX()
	friendsInfo, err := tx.GetFriendsShowsOrderByShowTime(accountID)
	if err != nil {
		zap.S().Infof("tx.GetShowsOrderByShowTime FAILED: %v", err)
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	groupInfo, err := tx.GetGroupShowsOrderBy(accountID)
	if err != nil {
		zap.S().Infof("tx.GetShowsOrderByShowTime FAILED: %v", err)
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}

	var totalSettingReq reply.SettingReq

	//两个pinTime都是按照降序排列的，所以直接使用双指针进行排序
	var i, j int
	for i, j = 0, 0; i < int(friendsInfo.Total) && j < int(groupInfo.Total); {
		if friendsInfo.Data[i].BaseSetting.PinTime.After(*groupInfo.Data[j].BaseSetting.PinTime) {
			totalSettingReq.Data = append(totalSettingReq.Data, friendsInfo.Data[i])
			totalSettingReq.Total++
			i++
		} else {
			totalSettingReq.Data = append(totalSettingReq.Data, groupInfo.Data[j])
			totalSettingReq.Total++
			j++
		}
	}
	for i < int(friendsInfo.Total) {
		totalSettingReq.Data = append(totalSettingReq.Data, friendsInfo.Data[i])
		totalSettingReq.Total++
		i++
	}
	for j < int(groupInfo.Total) {
		totalSettingReq.Data = append(totalSettingReq.Data, groupInfo.Data[j])
		totalSettingReq.Total++
		j++
	}
	totalSettingReq.Total = friendsInfo.Total + groupInfo.Total
	return &totalSettingReq, nil
}

func (setting) UpdatePins(ctx *gin.Context, accountID uint, isPins bool, relationID int64) errcode.Err {

	//先去判断一下relationID是否是存在的
	qS := query.NewSetting()
	settingInfo, ok := qS.CheckRelationIDExist(int64(accountID), relationID)

	if !ok {
		return myerr.DoNotHaveThisRelation
	}
	if settingInfo.IsPin == isPins {
		return nil
	}
	tokenString := ctx.GetHeader(global.Settings.Token.AuthType)
	if !ok {
		return errcode.ErrUnauthorizedAuthNotExist
	}
	if err := qS.UpdatePinsState(relationID, isPins); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	global.Worker.SendTask(task.UpdateSettingState(tokenString, serve.SettingPin, int64(settingInfo.RelationID)))
	return nil
}

func (setting) UpdateNickName(ctx *gin.Context, accountID uint, req request.UpdateNickName) errcode.Err {
	//修改关系线对应的昵称
	//注意这里只是修改了好友或者是群的备注，并没有修改本身的nickname
	qS := query.NewSetting()
	settingInfo, ok := qS.CheckRelationIDExist(int64(accountID), req.RelationID)

	if !ok {
		return myerr.DoNotHaveThisRelation
	}
	if settingInfo.NickName == req.NickName {
		return nil
	}
	if err := qS.UpdateNickName(req.NickName, req.RelationID); err != nil {
		return errcode.ErrServer
	}
	tokenString := ctx.GetHeader(global.Settings.Token.AuthType)
	global.Worker.SendTask(task.UpdateSettingNickName(tokenString, req.RelationID))
	return nil
}

func (setting) UpdateDisturbState(ctx *gin.Context, accountID uint, req request.UpdateIsDisturbState) errcode.Err {
	qS := query.NewSetting()
	settingInfo, ok := qS.CheckRelationIDExist(int64(accountID), req.RelationID)

	if !ok {
		return myerr.DoNotHaveThisRelation
	}
	if settingInfo.IsNotDisturbed == req.IsDisturbState {
		return nil
	}
	if err := qS.UpdateIsDisturbState(req.RelationID, req.IsDisturbState); err != nil {
		zap.S().Infof("UpdateIsDisturbState failed err: %v", err)
		return errcode.ErrServer.WithDetails(err.Error())
	}
	tokenString := ctx.GetHeader(global.Settings.Token.AuthType)
	global.Worker.SendTask(task.UpdateSettingIsDisturbState(tokenString, req.RelationID, req.IsDisturbState))
	return nil
}

func (setting) UpdateShowState(ctx *gin.Context, accountID uint, req request.UpdateShowState) errcode.Err {
	qS := query.NewSetting()
	settingInfo, ok := qS.CheckRelationIDExist(int64(accountID), req.RelationID)
	if !ok {
		return myerr.DoNotHaveThisRelation
	}
	if settingInfo.IsShow == req.IsShow {
		return nil
	}
	if err := qS.UpdateIsShowState(req.RelationID, req.IsShow); err != nil {
		zap.S().Infof("UpdateIsDisturbState failed err: %v", err)
		return errcode.ErrServer.WithDetails(err.Error())
	}
	tokenString := ctx.GetHeader(global.Settings.Token.AuthType)
	global.Worker.SendTask(task.UpdateSettingShow(tokenString, req.RelationID, req.IsShow))
	return nil
}
