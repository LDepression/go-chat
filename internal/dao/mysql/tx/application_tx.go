package tx

import (
	"errors"
	"go-chat/internal/dao"
	"go-chat/internal/dao/mysql/query"
	"go-chat/internal/model/automigrate"
	"go-chat/internal/model/common"
	"go-chat/internal/model/request"
	"gorm.io/gorm"
)

var (
	ErrHasThisFriend = errors.New("已经是好友了")
	ErrIsSelf        = errors.New("不能添加自己为好友")
)

type applicationTX struct {
}

func NewApplicationTX() *applicationTX {
	return &applicationTX{}
}

// CreateApplicationWithTX 第一个参数是申请者,第二个参数是被申请者
func (applicationTX) CreateApplicationWithTX(account1ID, account2ID uint64, ApplyMsg string) (uint64, error) {
	//先去判断一下目标id是不是自己
	if account2ID == account1ID {
		return 0, ErrIsSelf
	}
	//先去判断一下两者是否已经是好友了
	qR := query.NewRelation()
	if ok := qR.CheckISFriend(account1ID, account2ID); ok {
		return 0, ErrHasThisFriend
	}
	q := query.NewApplication()
	applicationID, err := q.CreateApplication(account1ID, request.CreateApplicationReq{
		AccountID:      account2ID,
		ApplicationMsg: ApplyMsg,
	})
	if err != nil {
		return 0, err
	}
	return applicationID, nil
}
func (applicationTX) AcceptApplicationWithTX(applicantID, receiverID uint) error {

	tx := dao.Group.DB.Begin()
	defer tx.Commit()
	// 将application中status更新为"已接受"
	if result := tx.Model(&automigrate.Application{}).
		Where("account_id1 = ? AND account_id2 = ?", applicantID, receiverID).
		Update("status", common.ApplicationStateAccepted); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// 建立两个人之间的relation
	var relation automigrate.Relation
	relation.RelationType = "friend"
	relation.FriendType = automigrate.FriendType{
		AccountID1: int64(applicantID),
		AccountID2: int64(receiverID),
	}
	if result := tx.Model(&automigrate.Relation{}).Create(&relation); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// 对两个account分别设置setting
	var settingAcceptInfo, settingReceiverInfo automigrate.Setting
	settingAcceptInfo.RelationID, settingReceiverInfo.RelationID = relation.ID, relation.ID
	//	增添对方id的setting
	settingAcceptInfo.AccountID, settingReceiverInfo.AccountID = receiverID, applicantID
	if result := tx.Model(&automigrate.Setting{}).Create(&settingAcceptInfo).Create(&settingReceiverInfo); result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return nil
}

func CommitTXFailed(tx *gorm.DB) error {
	if r := recover(); r != nil {
		tx.Rollback()
	}
	if err := tx.Commit().Error; err != nil {
		// 处理事务提交失败的错误
		return err
	}
	return nil
}
