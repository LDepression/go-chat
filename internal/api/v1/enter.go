/**
 * @Author: lenovo
 * @Description:
 * @File:  enter
 * @Version: 1.0.0
 * @Date: 2023/03/20 22:54
 */

package v1

import (
	"go-chat/internal/api/v1/chat"
)

type group struct {
	User        user
	Email       email
	Account     account
	Application application
	File        file
	Setting     setting
	Chat        chat.Group
	VGroup      vgroup
}

var Group = new(group)
