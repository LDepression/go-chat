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
	worker, _ := NewWorker(1)
	id := worker.GetId()
	fmt.Println(id)
}
