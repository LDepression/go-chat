/**
 * @Author: lenovo
 * @Description:
 * @File:  messsage
 * @Version: 1.0.0
 * @Date: 2023/05/14 22:54
 */

package query

import (
	"go-chat/internal/dao"
	"go-chat/internal/model/automigrate"
	"gorm.io/gorm"
)

type message struct{}

func NewMessage() *message {
	return &message{}
}

func (message) CheckMsgInfoByID(msgID int64) (*automigrate.Message, error) {
	var msgInfo automigrate.Message
	if result := dao.Group.DB.Model(&automigrate.Message{}).Where("id = ?", msgID).Find(&msgInfo); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &msgInfo, nil
}

func (message) CreateMsg(params CreateMsgParams) (*automigrate.Message, error) {
	Msg := &automigrate.Message{
		MsgType:    automigrate.MsgType(params.MsgType),
		MsgContent: params.MsgContent,
		MsgExtend:  params.MsgExtend,
		AccountID:  params.AccountID,
		RlyMsgID:   params.ReplyMsgID,
		RelationID: params.RelationID.Int64,
		IsRevoke:   false,
		IsTop:      false,
		IsPin:      false,
	}
	if result := dao.Group.DB.Model(&automigrate.Message{}).Create(Msg); result.RowsAffected == 0 {
		return nil, result.Error
	}
	return Msg, nil
}
