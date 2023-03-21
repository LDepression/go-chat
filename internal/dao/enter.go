/**
 * @Author: lenovo
 * @Description:
 * @File:  enter
 * @Version: 1.0.0
 * @Date: 2023/03/20 15:27
 */

package dao

import (
	"go-chat/internal/dao/redis/query"
	"gorm.io/gorm"
)

type group struct {
	DB    *gorm.DB
	Redis *query.Queries
}

var Group = new(group)
