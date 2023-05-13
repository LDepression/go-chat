/**
 * @Author: lenovo
 * @Description:
 * @File:  group_test
 * @Version: 1.0.0
 * @Date: 2023/05/07 20:21
 */

package query

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-chat/internal/model/request"
	"testing"
)

func TestGroup_CreateGroupRelation(t *testing.T) {
	qGroup := NewGroup()
	Rid, err := qGroup.CreateGroupRelation(request.CreateGroupReq{
		SigNature: "剑指offer",
		Name:      "gopher",
	})
	assert.NoError(t, err)
	//assert.Equal(t, 12, Rid)
	fmt.Println(Rid)
}
