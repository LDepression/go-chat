package query

import (
	"fmt"
	"go-chat/internal/dao"
	"testing"
)

func TestGetApplicationByID(t *testing.T) {
	InitMySql()
	if dao.Group.DB == nil {
		fmt.Println("dao.Group.DB == nil")
	}
	qApplication := NewQueryApplication()
	applicationInfo, err := qApplication.GetApplicationByID(7041606952899575808, 2046824448)
	if err != nil {
		fmt.Printf("\n\tapplicationInfo, err := , err:%#v \n", err)
		t.Errorf("err:%v\n,", err)
	}
	fmt.Println(applicationInfo)
}

func TestRefuseApplication(t *testing.T) {
	InitMySql()
	if dao.Group.DB == nil {
		fmt.Println("dao.Group.DB == nil")
	}
	qApplication := NewQueryApplication()
	err := qApplication.RefuseApplication(7041606952899575808, 7042309196552863744, "")
	if err != nil {
		fmt.Printf("\n\tapplicationInfo, err := , err:%#v \n", err)
		t.Errorf("err:%v\n,", err)
	}
}

func TestGetApplicationList(t *testing.T) {
	InitMySql()
	if dao.Group.DB == nil {
		fmt.Println("dao.Group.DB == nil")
	}
	qApplication := NewQueryApplication()
	Info, total, err := qApplication.GetApplicationsList(7042309196552863744, 10, 0)
	if err != nil {
		fmt.Printf("\n\tapplicationInfo, err := , err:%#v \n", err)
		t.Errorf("err:%v\n,", err)
	}
	for _, v := range Info {
		fmt.Println(v)
	}
	fmt.Println("total:", total)
}
