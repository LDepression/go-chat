/**
 * @Author: lenovo
 * @Description:
 * @File:  setting
 * @Version: 1.0.0
 * @Date: 2023/05/05 20:48
 */

package task

import (
	"go-chat/internal/global"
	"go-chat/internal/middleware"
	"go-chat/internal/model"
	"go-chat/internal/model/chat"
	"go-chat/internal/model/chat/serve"
	"go-chat/internal/pkg/utils"
	"go.uber.org/zap"
)

func UpdateSettingState(accessToken string, settingType serve.SettingType, relationID int64) func() {
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
		global.ChatMap.Send(int64(content.ID), chat.ServerUpdateAccount, serve.UpdateSettingType{
			Encode:     utils.EncodeMD5(token),
			RelationID: relationID,
			SType:      settingType,
		})
	}
}

func UpdateSettingNickName(accessToken string, relationID int64) func() {
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
		global.ChatMap.Send(int64(content.ID), chat.ServerUpdateAccount, serve.UpdateNickName{
			Encode:     utils.EncodeMD5(token),
			RelationID: relationID,
		})
	}
}

func UpdateSettingIsDisturbState(accessToken string, relationID int64, isDisturbState bool) func() {
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
		global.ChatMap.Send(int64(content.ID), chat.ServerUpdateAccount, serve.UpdateIsDisturbState{
			Encode:         utils.EncodeMD5(token),
			RelationID:     relationID,
			SType:          serve.SettingNotDisturb,
			IsDisturbState: isDisturbState,
		})

	}
}

func UpdateSettingShow(accessToken string, relationID int64, isShowState bool) func() {
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
		global.ChatMap.Send(int64(content.ID), chat.ServerUpdateAccount, serve.UpdateShowState{
			Encode:     utils.EncodeMD5(token),
			RelationID: relationID,
			SType:      serve.SettingNotDisturb,
			IsShow:     isShowState,
		})
	}
}
