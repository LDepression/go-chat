package query

//
//import (
//	"fmt"
//	"go-chat/internal/dao"
//	"testing"
//)
//
//func TestGetRelationInfoByAccountsID(t *testing.T) {
//	InitMySql()
//	if dao.Group.DB == nil {
//		fmt.Println("dao.Group.DB == nil")
//	}
//	qSetting := NewQuerySetting()
//	relationInfo, err := qSetting.GetRelationInfoByAccountsID(2, 1)
//	if err != nil {
//		fmt.Printf("qSetting.GetRelationInfoByAccountsID, err:%#v \n", err)
//		t.Errorf("err:%v\n,", err)
//	}
//	fmt.Println(relationInfo)
//}
//
//func TestGetFriendInfoByID(t *testing.T) {
//	InitMySql()
//	if dao.Group.DB == nil {
//		fmt.Println("dao.Group.DB == nil")
//	}
//	qSetting := NewQuerySetting()
//	FriendInfo, err := qSetting.GetFriendInfoByID(62, 7044171529629728768)
//	if err != nil {
//		fmt.Printf("qSetting.GetFriendInfo failed, err:%#v \n", err)
//		t.Errorf("err:%v\n,", err)
//	}
//	fmt.Println(FriendInfo)
//	fmt.Println("Relation:", FriendInfo.Relation)
//	fmt.Println("Account:", FriendInfo.Account)
//}
//
//func TestGetFriendInfoByName(t *testing.T) {
//	InitMySql()
//	if dao.Group.DB == nil {
//		fmt.Println("dao.Group.DB == nil")
//	}
//	qSetting := NewQuerySetting()
//	FriendInfo, err := qSetting.GetFriendInfoByName(7044171305423208448, 61, 10, 0, "why")
//	if err != nil {
//		fmt.Printf("qSetting.GetFriendInfoByName, err:%#v \n", err)
//		t.Errorf("err:%v\n,", err)
//		return
//	}
//	fmt.Println(FriendInfo)
//	fmt.Println("Relation:", FriendInfo.Relation)
//	fmt.Println("Account:", FriendInfo.Account)
//}
//
//func TestGetRelationInfoByRelationID(t *testing.T) {
//	InitMySql()
//	if dao.Group.DB == nil {
//		fmt.Println("dao.Group.DB == nil")
//	}
//		t.Errorf("err:%v\n,", err)
//	}
//	fmt.Println(relationInfo)
//}
//
//func TestUpdateNickName(t *testing.T) {
//	InitMySql()
//	if dao.Group.DB == nil {
//		fmt.Println("dao.Group.DB == nil")
//	}
//	qSetting := NewQuerySetting()
//	err := qSetting.UpdateNickName(7044171529629728768, 62, "asdasd")
//	if err != nil {
//		fmt.Printf("qSetting.UpdateNickName, err:%#v \n", err)
//		t.Errorf("err:%v\n,", err)
//	}
//}
