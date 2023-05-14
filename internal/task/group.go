/**
 * @Author: lenovo
 * @Description:
 * @File:  group
 * @Version: 1.0.0
 * @Date: 2023/05/08 20:58
 */

package task

import (
	"context"
	"go-chat/internal/dao"
	"go-chat/internal/global"
	"go-chat/internal/middleware"
	"go-chat/internal/model"
	"go-chat/internal/model/chat"
	"go-chat/internal/model/chat/serve"
	"go-chat/internal/pkg/utils"
	"go.uber.org/zap"
	"time"
)

func SendInvitedInfoToInviters(accessToken string, InviterIDs []int64) func() {
	return func() {
		var content model.Content
		payload, token, err := middleware.ParseHeader(accessToken)
		if err != nil {
			return
		}
		if err := content.UnMarshal(payload.Content); err != nil {
			zap.S().Infof("unmarshal content: %v", err)
			return
		}
		global.ChatMap.SendMany(InviterIDs, chat.ServerInviteAccount, serve.InviteNewPerson{
			Encoding:  utils.EncodeMD5(token),
			AccountID: int64(content.ID),
		})
	}
}

func DissolveGroup(accessToken string, relationID int64) func() {
	var content model.Content
	payload, _, err := middleware.ParseHeader(accessToken)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	accountIDs, err1 := dao.Group.Redis.GetAllAccountsByRelationID(ctx, relationID)
	if err1 != nil {
		zap.S().Infof("dao.Group.Redis.GetAllAccountsByRelationID failed: %v", err)
		return func() {}

	}
	if err := content.UnMarshal(payload.Content); err != nil {
		zap.S().Infof("unmarshal content: %v", err)
		return func() {}

	}
	return func() {
		global.ChatMap.SendMany(accountIDs, chat.ServerGroupDissolved, serve.DissolveGroup{
			Encoding:  utils.EncodeMD5(accessToken),
			AccountID: int64(content.ID),
		})
	}
}

func TransferGroup(accessToken string, relationID int64) func() {
	var content model.Content
	payload, _, err := middleware.ParseHeader(accessToken)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	accountIDs, err1 := dao.Group.Redis.GetAllAccountsByRelationID(ctx, relationID)
	if err1 != nil {
		zap.S().Infof("dao.Group.Redis.GetAllAccountsByRelationID failed: %v", err)
		return func() {}

	}
	if err := content.UnMarshal(payload.Content); err != nil {
		zap.S().Infof("unmarshal content: %v", err)
		return func() {}
	}
	return func() {
		global.ChatMap.SendMany(accountIDs, chat.ServerGroupTransferred, serve.TransferGroup{
			Encoding:  utils.EncodeMD5(accessToken),
			AccountID: int64(content.ID),
		})
	}
}

func QuitGroup(accessToken string, relationID int64) func() {
	var content model.Content
	payload, _, err := middleware.ParseHeader(accessToken)
	if err != nil {
		zap.S().Infof("parse header: %v", err)
		return func() {}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	members, err1 := dao.Group.Redis.GetAllAccountsByRelationID(ctx, relationID)
	if err1 != nil {
		zap.S().Infof("error : %v", err1)
		return func() {}
	}
	if err := content.UnMarshal(payload.Content); err != nil {
		zap.S().Infof("unmarshal content: %v", err)
		return func() {}
	}
	return func() {
		global.ChatMap.SendMany(members, chat.ServerQuitGroup, serve.QuitGroup{
			Encoding:  utils.EncodeMD5(accessToken),
			AccountID: int64(content.ID),
		})
	}
}
