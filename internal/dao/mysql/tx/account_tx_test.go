/**
 * @Author: lenovo
 * @Description:
 * @File:  account_tx_test
 * @Version: 1.0.0
 * @Date: 2023/03/29 18:50
 */

package tx

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

//func TestCreateAccountWithTX(t *testing.T) {
//	tx := NewAccountTX()
//	ctx := context.Background()
//	err := tx.CreateAccountWithTX(ctx, 1, request.CreateAccountReq{
//		ID:        global.SnowFlake.GetId(),
//		Name:      "shoyich2",
//		Signature: "golang yyds",
//		Gender:    0,
//	})
//	require.NoError(t, err)
//}

func TestDeleteAccountWithTX(t *testing.T) {
	tx := NewAccountTX()
	ctx := context.Background()
	err := tx.DeleteAccountWithTX(ctx, 3)
	require.NoError(t, err)
}
