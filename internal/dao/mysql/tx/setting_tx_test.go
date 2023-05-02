package tx

import (
	"fmt"
	"go-chat/internal/dao"
	"testing"
)

func TestDeleteFriendWithTX(t *testing.T) {
	InitMySql()
	if dao.Group.DB == nil {
		fmt.Println("dao.Group.DB == nil")
	}
	settingTX := NewSettingTX()
	err := settingTX.DeleteFriendWithTX(60)
	if err != nil {
		fmt.Println("为啥啊？", err)
		fmt.Printf("\n\tapplicationInfo, err := , err:%#v \n", err)
		t.Errorf("err:%v\n,", err)
	}
}
