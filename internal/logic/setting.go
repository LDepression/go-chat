package logic

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/dao/mysql/query"
	"go-chat/internal/dao/mysql/tx"
	"go-chat/internal/model/automigrate"
	"go-chat/internal/model/common"
	"go-chat/internal/model/reply"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app/errcode"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type setting struct {
}

func (setting) DeleteFriend(c *gin.Context, selfAccountID, targetAccountID uint) errcode.Err {
	qSetting := query.NewQuerySetting()
	// 通过对方accountID获取两人之间的 relationID
	relationInfo, err := qSetting.GetRelationInfoByAccountsID(selfAccountID, targetAccountID)
	if err != nil {
		zap.S().Errorf("qSetting.GetRelationInfoByAccountsID failed,err:%v", err)
		if err == gorm.ErrRecordNotFound {
			return myerr.RelationNotExist
		}
		return errcode.ErrServer
	}
	// 通过relationID删除 relation和settings
	settingTX := tx.NewSettingTX()
	err = settingTX.DeleteFriendWithTX(relationInfo.ID)
	if err != nil {
		zap.S().Errorf("settingTX.DeleteFriendWithTX failed,err:%v", err)
		return errcode.ErrServer
	}
	return nil
}

func (setting) GetFriendsList(c *gin.Context, selfAccountID uint) (*reply.GetFriendsList, errcode.Err) {
	qSetting := query.NewQuerySetting()
	RelationInfos, err := qSetting.GetRelationInfos(selfAccountID)
	if err != nil {
		zap.S().Errorf("qSetting.GetFriendsAccounts failed,err:%v", err)
		return &reply.GetFriendsList{}, errcode.ErrServer
	}
	if len(RelationInfos) == 0 {
		return &reply.GetFriendsList{}, nil
	}
	FriendsInfos, err := GetFriendsInfosByID(selfAccountID, RelationInfos)
	if err != nil {
		zap.S().Errorf("GetFriendsInfosByID failed,err:%v", err)
		return &reply.GetFriendsList{}, errcode.ErrServer
	}
	return &reply.GetFriendsList{
		FriendsInfos: FriendsInfos,
		Total:        len(FriendsInfos),
	}, nil
}

func GetFriendsInfosByID(selfAccountID uint, relationInfos []*automigrate.Relation) ([]*reply.GetFriendInfo, error) {
	qSetting := query.NewQuerySetting()
	FriendsInfos := make([]*reply.GetFriendInfo, 0, len(relationInfos))
	for _, v := range relationInfos {
		relationID := v.ID
		var targetAccountID uint
		var FriendInfo *automigrate.Setting
		var err error
		if uint(v.FriendType.AccountID1) == selfAccountID {
			targetAccountID = uint(v.FriendType.AccountID2)
			FriendInfo, err = qSetting.GetFriendInfoByID(targetAccountID, relationID)
			if err != nil {
				return nil, err
			}
		} else if uint(v.FriendType.AccountID2) == selfAccountID {
			targetAccountID = uint(v.FriendType.AccountID1)
			FriendInfo, err = qSetting.GetFriendInfoByID(targetAccountID, relationID)
			if err != nil {
				return nil, err
			}
		}
		if FriendInfo == nil {
			return nil, nil
		}
		FriendsInfos = append(FriendsInfos, &reply.GetFriendInfo{
			Account: reply.EasyAccount{
				AccountID: targetAccountID,
				Name:      FriendInfo.Account.Name,
				Avatar:    FriendInfo.Account.Avatar,
			},
			Setting: reply.EasySetting{
				RelationID:     relationID,
				RelationType:   common.RelationTypeFriend,
				NickName:       FriendInfo.NickName,
				IsNotDisturbed: FriendInfo.IsNotDisturbed,
				IsPin:          FriendInfo.IsPin,
				PinTime:        FriendInfo.PinTime,
				IsShow:         FriendInfo.IsShow,
				LastShowTime:   FriendInfo.LastShowTime,
				IsSelf:         FriendInfo.IsSelf,
			},
		})
	}
	return FriendsInfos, nil
}

func (setting) GetFriendsByName(c *gin.Context, selfAccountID, limit, offset uint, name string) (*reply.GetFriendsByName, errcode.Err) {
	if name == "" {
		return &reply.GetFriendsByName{}, nil
	}
	qSetting := query.NewQuerySetting()
	RelationInfos, err := qSetting.GetRelationInfos(selfAccountID)
	if err != nil {
		zap.S().Errorf("qSetting.GetFriendsAccounts failed,err:%v", err)
		return &reply.GetFriendsByName{}, errcode.ErrServer
	}
	if len(RelationInfos) == 0 {
		zap.S().Infof("relationInfos = %v", RelationInfos)
		return &reply.GetFriendsByName{}, nil
	}
	FriendsInfos, err := GetFriendsInfosByName(selfAccountID, limit, offset, name, RelationInfos)
	if err != nil {
		zap.S().Errorf("GetFriendsInfosByName failed,err:%v", err)
		return &reply.GetFriendsByName{}, errcode.ErrServer
	}
	return &reply.GetFriendsByName{
		FriendsInfos: FriendsInfos,
		Total:        len(FriendsInfos),
	}, nil
}

func GetFriendsInfosByName(selfAccountID, limit, offset uint, name string, relationInfos []*automigrate.Relation) ([]*reply.GetFriendInfo, error) {
	qSetting := query.NewQuerySetting()
	FriendsInfos := make([]*reply.GetFriendInfo, 0, len(relationInfos))
	for _, v := range relationInfos {
		relationID := v.ID
		var targetAccountID uint
		var FriendInfo *automigrate.Setting
		var err error
		if uint(v.FriendType.AccountID1) == selfAccountID {
			targetAccountID = uint(v.FriendType.AccountID2)
			FriendInfo, err = qSetting.GetFriendInfoByName(targetAccountID, relationID, limit, offset, name)
			if err != nil {
				return nil, err
			}
		} else if uint(v.FriendType.AccountID2) == selfAccountID {
			targetAccountID = uint(v.FriendType.AccountID1)
			FriendInfo, err = qSetting.GetFriendInfoByName(targetAccountID, relationID, limit, offset, name)
			if err != nil {
				return nil, err
			}
		}
		if FriendInfo == nil {
			return nil, nil
		}
		FriendsInfos = append(FriendsInfos, &reply.GetFriendInfo{
			Account: reply.EasyAccount{
				AccountID: targetAccountID,
				Name:      FriendInfo.Account.Name,
				Avatar:    FriendInfo.Account.Avatar,
			},
			Setting: reply.EasySetting{
				RelationID:     relationID,
				RelationType:   common.RelationTypeFriend,
				NickName:       FriendInfo.NickName,
				IsNotDisturbed: FriendInfo.IsNotDisturbed,
				IsPin:          FriendInfo.IsPin,
				PinTime:        FriendInfo.PinTime,
				IsShow:         FriendInfo.IsShow,
				LastShowTime:   FriendInfo.LastShowTime,
				IsSelf:         FriendInfo.IsSelf,
			},
		})
	}
	return FriendsInfos, nil
}

func (setting) UpdateNickName(c *gin.Context, selfAccountID, relationID uint, nickName string) errcode.Err {
	// 通过relationID获取relationInfo
	qSetting := query.NewQuerySetting()
	relationInfo, err := qSetting.GetRelationInfoByRelationID(relationID)
	if err != nil {
		zap.S().Errorf("qSetting.GetRelationInfoByRelationID failed,err:%v", err)
		if err == gorm.ErrRecordNotFound {
			return myerr.RelationNotExist
		}
		return errcode.ErrServer
	}
	// 通过relationType判断修改的信息
	switch relationInfo.RelationType {
	case common.RelationTypeFriend:
		if uint(relationInfo.FriendType.AccountID1) == selfAccountID {
			targetAccountID := uint(relationInfo.FriendType.AccountID2)
			err = qSetting.UpdateNickName(targetAccountID, relationID, nickName)
		} else if uint(relationInfo.FriendType.AccountID2) == selfAccountID {
			targetAccountID := uint(relationInfo.FriendType.AccountID1)
			err = qSetting.UpdateNickName(targetAccountID, relationID, nickName)
		}
		if err != nil {
			zap.S().Errorf("qSetting.UpdateNickName failed,err:%v", err)
			return errcode.ErrServer
		}
	case common.RelationTypeGroup:
		err = qSetting.UpdateNickName(0, relationID, nickName)
		if err != nil {
			zap.S().Errorf("qSetting.UpdateNickName failed,err:%v", err)
			return errcode.ErrServer
		}
	default:
		return errcode.ErrNotFound
	}
	return nil
}

func (setting) UpdateSettingDisturb(c *gin.Context, selfAccountID, relationID uint, isDisturbed bool) errcode.Err {
	// 通过relationID获取relationInfo
	qSetting := query.NewQuerySetting()
	relationInfo, err := qSetting.GetRelationInfoByRelationID(relationID)
	if err != nil {
		zap.S().Errorf("qSetting.GetRelationInfoByRelationID failed,err:%v", err)
		if err == gorm.ErrRecordNotFound {
			return myerr.RelationNotExist
		}
		return errcode.ErrServer
	}
	// 通过relationType判断修改的信息
	switch relationInfo.RelationType {
	case common.RelationTypeFriend:
		if uint(relationInfo.FriendType.AccountID1) == selfAccountID {
			targetAccountID := uint(relationInfo.FriendType.AccountID2)
			err = qSetting.UpdateSettingDisturb(targetAccountID, relationID, isDisturbed)
		} else if uint(relationInfo.FriendType.AccountID2) == selfAccountID {
			targetAccountID := uint(relationInfo.FriendType.AccountID1)
			err = qSetting.UpdateSettingDisturb(targetAccountID, relationID, isDisturbed)
		}
		if err != nil {
			zap.S().Errorf("qSetting.UpdateNickName failed,err:%v", err)
			return errcode.ErrServer
		}
	case common.RelationTypeGroup:
		err = qSetting.UpdateSettingDisturb(0, relationID, isDisturbed)
		if err != nil {
			zap.S().Errorf("qSetting.UpdateNickName failed,err:%v", err)
			return errcode.ErrServer
		}
	default:
		return errcode.ErrNotFound
	}
	return nil
}
