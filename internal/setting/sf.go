/**
 * @Author: lenovo
 * @Description:
 * @File:  sf
 * @Version: 1.0.0
 * @Date: 2023/03/29 21:37
 */

package setting

import (
	"go-chat/internal/global"
	"go-chat/internal/pkg/snowflake"
)

type sf struct {
}

func (sf) Init() {
	global.SnowFlake, _ = snowflake.NewWorker(1)
}
