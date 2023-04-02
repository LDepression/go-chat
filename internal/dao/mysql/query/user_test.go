/**
 * @Author: lenovo
 * @Description:
 * @File:  user_test
 * @Version: 1.0.0
 * @Date: 2023/03/23 22:35
 */

package query

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-chat/internal/model/automigrate"
	"testing"
)

func TestSaveRegisterInfo(t *testing.T) {
	quser := NewQueryUser()
	userID, err := quser.SaveRegisterInfo(automigrate.User{
		Email:    "1197285121@qq.com",
		Mobile:   "13877777776",
		Password: "123456",
	})
	assert.NoError(t, err)
	fmt.Println(userID)
}

func TestGetUserByID(t *testing.T) {
	quser := NewQueryUser()
	userInfo, err := quser.GetUserByID(300)
	assert.NoError(t, err)
	fmt.Println(*userInfo)
}
