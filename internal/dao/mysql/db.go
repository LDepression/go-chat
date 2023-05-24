/**
 * @Author: lenovo
 * @Description:
 * @File:  db.go
 * @Version: 1.0.0
 * @Date: 2023/03/20 15:06
 */

package mysql

import (
	"fmt"
	"go-chat/internal/dao"
	"go-chat/internal/global"
	"go-chat/internal/model/automigrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func InitMySql() {
	m := global.Settings.Mysql
	fmt.Println(m)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.User, m.Password, m.Host, m.Port, m.DbName)
	fmt.Println(dsn)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢阙值
			Colorful:      true,        //禁用彩色
			LogLevel:      logger.Info,
		})
	//全局模式
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if err != nil {
		panic(err)
	}
	dao.Group.DB = DB
	_ = DB.AutoMigrate(&automigrate.User{}, &automigrate.Account{}, &automigrate.Relation{}, &automigrate.Setting{}, &automigrate.Application{}, &automigrate.File{}, &automigrate.Notify{})
	fmt.Println("数据库连接成功！！！")
}
