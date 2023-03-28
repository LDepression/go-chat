/**
 * @Author: lenovo
 * @Description:
 * @File:  account
 * @Version: 1.0.0
 * @Date: 2023/03/28 17:27
 */

package logic

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/model/request"
)

type account struct{}

func (account) CreateAccount(ctx *gin.Context, req request.CreateAccountReq) {

}
