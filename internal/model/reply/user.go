/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/03/20 20:22
 */

package reply

import (
	"go-chat/internal/pkg/token"
)

type email struct{}

type LoginReply struct {
	UserID       int64         `json:"UserID"`
	AccessToken  string        `json:"AccessToken"`
	RefreshToken string        `json:"RefreshToken"`
	PayLoad      token.Payload `json:"Payload"`
}
