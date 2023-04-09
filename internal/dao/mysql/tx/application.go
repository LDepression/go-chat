/**
 * @Author: lenovo
 * @Description:
 * @File:  application
 * @Version: 1.0.0
 * @Date: 2023/04/04 9:28
 */

package tx

import (
	"errors"
	query "go-chat/internal/dao/mysql/query"
	"go-chat/internal/model/request"
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
	if uint64(account2ID) == account1ID {
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
