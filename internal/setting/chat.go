/**
 * @Author: lenovo
 * @Description:
 * @File:  chat
 * @Version: 1.0.0
 * @Date: 2023/05/01 1:07
 */

package setting

import (
	"go-chat/internal/global"
	"go-chat/internal/manager"
)

type chat struct{}

func (chat) Init() {
	global.ChatMap = manager.NewChatMap()
}
