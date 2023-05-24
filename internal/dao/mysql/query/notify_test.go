package query

import (
	"fmt"
	"go-chat/internal/dao"
	"go-chat/internal/model/automigrate"
	"gorm.io/gorm"
	"testing"
)

func TestCreateNotify(t *testing.T) {
	InitMySql()
	if dao.Group.DB == nil {
		fmt.Println("dao.Group.DB == nil")
	}
	qNotify := NewQueryNotify()
	params := automigrate.Notify{
		RelationID: 999,
		AccountID:  999,
		MsgContent: "这是999条公告",
		MsgExpand:  &automigrate.MsgExpand{Reminds: []automigrate.Remind{{1999, 10000}, {2999, 20000}}},
	}
	notify, err := qNotify.CreateNotify(params.AccountID, params.RelationID, params.MsgContent, params.MsgExpand)
	if err != nil {
		fmt.Printf("qAccount.GetAccountByID failed, err:%#v \n", err)
		t.Errorf("err:%v\n,", err)
	}
	fmt.Println("notify.ID = ", notify.ID)
	fmt.Println(notify)
}

func TestCheckIsLeader(t *testing.T) {
	InitMySql()
	if dao.Group.DB == nil {
		fmt.Println("dao.Group.DB == nil")
	}
	qNotify := NewQueryNotify()
	ok, err := qNotify.CheckIsLeader(7045205516011700224, 1)
	if err != nil {
		fmt.Printf("qNotify.CheckIsLeader failed, err:%#v \n", err)
		t.Errorf("err:%v\n,", err)
	}
	fmt.Println(ok)
}

func TestDeleteNotify(t *testing.T) {
	InitMySql()
	if dao.Group.DB == nil {
		fmt.Println("dao.Group.DB == nil")
	}
	qNotify := NewQueryNotify()
	err := qNotify.DeleteNotify(2, 123)
	if err != nil {
		fmt.Printf("qNotify.DeleteNotify failed, err:%#v \n", err)
		t.Errorf("err:%v\n,", err)
	}
	fmt.Println("success")
}

func TestUpdateNotify(t *testing.T) {
	InitMySql()
	if dao.Group.DB == nil {
		fmt.Println("dao.Group.DB == nil")
	}
	qNotify := NewQueryNotify()
	params := automigrate.Notify{
		Model:      gorm.Model{ID: 7},
		RelationID: 19,
		MsgContent: "这是22222222222条公告",
		MsgExpand:  &automigrate.MsgExpand{Reminds: []automigrate.Remind{{999123, 10000}, {222, 20000}}},
	}
	notify, err := qNotify.UpdateNotify(params.ID, params.RelationID, params.MsgContent, params.MsgExpand)
	if err != nil {
		fmt.Printf("qNotify.UpdateNotify failed, err:%#v \n", err)
		t.Errorf("err:%v\n,", err)
	}
	//fmt.Println("notify.ID = ", notify.ID)
	fmt.Println(notify)
}

func TestCheckIsInGroup(t *testing.T) {
	InitMySql()
	if dao.Group.DB == nil {
		fmt.Println("dao.Group.DB == nil")
	}
	qNotify := NewQueryNotify()
	ok, err := qNotify.CheckIsInGroup(7060068379461156864, 19)
	if err != nil {
		fmt.Printf("qNotify.CheckIsLeader failed, err:%#v \n", err)
		t.Errorf("err:%v\n,", err)
	}
	fmt.Println(ok)
}

func TestGetNotifies(t *testing.T) {
	InitMySql()
	if dao.Group.DB == nil {
		fmt.Println("dao.Group.DB == nil")
	}
	qNotify := NewQueryNotify()
	notifiesInfo, err := qNotify.GetNotifies(19)
	if err != nil {
		fmt.Printf("qNotify.CheckIsLeader failed, err:%#v \n", err)
		t.Errorf("err:%v\n,", err)
	}
	for _, v := range notifiesInfo {
		fmt.Printf("%#v\n", v)
	}
}
