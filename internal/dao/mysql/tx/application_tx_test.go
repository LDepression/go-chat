package tx

import (
	"fmt"
	"go-chat/internal/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func InitMySql() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/chat_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dao.Group.DB = db
}

func TestAcceptApplicationWithTX(t *testing.T) {
	InitMySql()
	if dao.Group.DB == nil {
		fmt.Println("dao.Group.DB == nil")
	}
	ApplicationTX := NewApplicationTX()
	err := ApplicationTX.AcceptApplicationWithTX(7041606952899575808, 7042309196552863744)
	if err != nil {
		fmt.Println("为啥啊？", err)
		fmt.Printf("\n\tapplicationInfo, err := , err:%#v \n", err)
		t.Errorf("err:%v\n,", err)
	}
}
