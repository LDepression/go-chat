/**
 * @Author: lenovo
 * @Description:
 * @File:  setting_tx_test
 * @Version: 1.0.0
 * @Date: 2023/04/18 21:48
 */

package tx

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func testSettingTX_GetFriendsPinsInfo(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{name: "goodCase", f: func() {
			tx := NewSettingTX()
			result, err := tx.GetFriendsPinsInfo(1837105152)
			require.NoError(t, err)
			for _, v := range result.Data {
				fmt.Println(v.Friend)
			}
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func testSettingTX_GetGroupsPinsInfo(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{name: "good", f: func() {
			tx := NewSettingTX()
			req, err := tx.GetGroupsPinsInfo(7046676613466947584)
			require.NoError(t, err)
			for _, v := range req.Data {
				//fmt.Println(v.BaseSetting.ID)
				fmt.Println(v.Group.Signature)
			}
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func TestTotal(t *testing.T) {
	testSettingTX_GetGroupsPinsInfo(t)
}
