/**
 * @Author: lenovo
 * @Description:
 * @File:  dao
 * @Version: 1.0.0
 * @Date: 2023/03/20 14:46
 */

package setting

import (
	"go-chat/internal/dao"
	"go-chat/internal/dao/mysql"
	"go-chat/internal/dao/redis"
)

type mdao struct {
}

func (d mdao) Init() {
	mysql.InitMySql()
	dao.Group.Redis = redis.Init()
}
