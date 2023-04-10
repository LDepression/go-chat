package query

import (
	"fmt"
	"go-chat/internal/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

//	func TestAddAccount(t *testing.T) {
//		accountInfo := &automigrate.Account{
//			Model: gorm.Model{
//				ID:        1,
//				CreatedAt: time.Now(),
//				UpdatedAt: time.Time{},
//				DeletedAt: gorm.DeletedAt{},
//			},
//			UserID:    3000,
//			User:      automigrate.User{},
//			Name:      "wang",
//			Signature: "person",
//			Avatar:    "",
//			Gender:    0,
//		}
//		if result := dao.Group.DB.Create(accountInfo); result.RowsAffected == 0 {
//			fmt.Printf("dao.Group.DB.Create() failed, err:%#v \n", result.Error)
//			t.Errorf("err:%v\n,", result.Error)
//		}
//	}
func InitMySql() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/chat_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dao.Group.DB = db
}

func TestGetAccountByID(t *testing.T) {
	InitMySql()
	if dao.Group.DB == nil {
		fmt.Println("dao.Group.DB == nil")
	}
	qAccount := NewQueryAccount()
	accountInfo, err := qAccount.GetAccountByID(1)
	if err != nil {
		fmt.Printf("qAccount.GetAccountByID failed, err:%#v \n", err)
		t.Errorf("err:%v\n,", err)
	}
	fmt.Println(accountInfo)
}

func TestGetAccountsByName(t *testing.T) {
	InitMySql()
	if dao.Group.DB == nil {
		fmt.Println("dao.Group.DB == nil")
	}
	qAccount := NewQueryAccount()
	accountInfos, total, err := qAccount.GetAccountsByName("wang", 1, 1)
	if err != nil {
		fmt.Printf("qAccount.GetAccountsByName failed, err:%#v \n", err)
		t.Errorf("err:%v\n,", err)
	}
	fmt.Println("total:", total)
	for _, v := range accountInfos {
		fmt.Println(v)
	}
}

func TestGetAccountsByUserID(t *testing.T) {
	InitMySql()
	if dao.Group.DB == nil {
		fmt.Println("dao.Group.DB == nil")
	}
	qAccount := NewQueryAccount()
	accountInfos, err := qAccount.GetAccountsByUserID(2)
	if err != nil {
		fmt.Printf("qAccount.GetAccountsByUserID failed, err:%#v \n", err)
		t.Errorf("err:%v\n,", err)
	}
	fmt.Println("total:", accountInfos)
	for _, v := range accountInfos {
		fmt.Println(v)
	}
}

func TestUpdateAccount(t *testing.T) {
	InitMySql()
	if dao.Group.DB == nil {
		fmt.Println("dao.Group.DB == nil")
	}
	qAccount := NewQueryAccount()
	err := qAccount.UpdateAccount(4, "wangda", "", "", "asd")
	if err != nil {
		fmt.Printf("qAccount.UpdateAccount failed, err:%#v \n", err)
		t.Errorf("err:%v\n,", err)
	}
}
