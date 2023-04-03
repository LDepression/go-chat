/**
 * @Author: lenovo
 * @Description:
 * @File:  redis
 * @Version: 1.0.0
 * @Date: 2023/03/20 20:34
 */

package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-chat/internal/dao/redis/query"
	"go-chat/internal/global"
	"go.uber.org/zap"
)

func Init() *query.Queries {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Settings.Redis.Addr,
		Password: global.Settings.Redis.Password, // 密码
		DB:       0,                              // 数据库
		PoolSize: global.Settings.Redis.PoolSize, // 连接池大小
	})
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		// Redis连接失败，进行相应处理
		zap.S().Errorf("redis init failed, err:%v", err)
		fmt.Println("redis初始化失败！！！！！")
		panic(err)
	}
	return query.New(rdb)
}
