/**
 * @Author: lenovo
 * @Description:
 * @File:  relation_test
 * @Version: 1.0.0
 * @Date: 2023/04/06 19:07
 */

package query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckISFriend(t *testing.T) {
	r := NewRelation()
	ok := r.CheckISFriend(7041587417215664128, 7041587417215664128)
	assert.True(t, ok)
}
