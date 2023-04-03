/**
 * @Author: lenovo
 * @Description:
 * @File:  gen_id_test
 * @Version: 1.0.0
 * @Date: 2023/03/29 21:02
 */

package snowflake

import (
	"fmt"
	"testing"
)

func TestSnowFlake(t *testing.T) {
	mySnow, _ := NewSnowFlake(0, 0) //生成雪花算法
	id, _ := mySnow.NextId()
	fmt.Println(id)
}
