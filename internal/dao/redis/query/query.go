/**
 * @Author: lenovo
 * @Description:
 * @File:  query
 * @Version: 1.0.0
 * @Date: 2023/03/20 20:35
 */

package query

import (
	"github.com/go-redis/redis/v8"
)

type Queries struct {
	rdb *redis.Client
}

func New(rdb *redis.Client) *Queries {
	return &Queries{rdb: rdb}
}
