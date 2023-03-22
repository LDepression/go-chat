/**
 * @Author: lenovo
 * @Description:
 * @File:  maker
 * @Version: 1.0.0
 * @Date: 2023/03/21 22:30
 */

package setting

import (
	"go-chat/internal/global"
	"go-chat/internal/pkg/token"
)

type maker struct {
}

func (maker) Init() {
	global.Maker, _ = token.NewPasetoMaker([]byte(global.Settings.Token.Key))
}
