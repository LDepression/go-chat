/**
 * @Author: lenovo
 * @Description:
 * @File:  handle
 * @Version: 1.0.0
 * @Date: 2023/05/05 22:10
 */

package task

import (
	"go-chat/internal/global"
	"go-chat/internal/model/chat"
	"go-chat/internal/model/chat/serve"
	"go-chat/internal/pkg/utils"
)

func AccountLogin(accessToken, address string, accountID int64) func() {
	return func() {
		global.ChatMap.Send(accountID, chat.ServerAccountLogin, serve.AccountLogin{
			EnToken: utils.EncodeMD5(accessToken),
			Address: address,
		})
	}
}
