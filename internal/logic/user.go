/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/03/20 20:07
 */

package logic

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app/errcode"
)

type user struct {
}

func (user) Register(ctx *gin.Context, mobile, email, password string) errcode.Err {
	//先判断一下邮箱是否已经存在
	exist, err := Group.Email.CheckEmailIsUsed(ctx, email)
	if err != nil {
		return myerr.EmailExists
	}
	//
}
