package v1

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/logic"
	"go-chat/internal/model/request"
	"go-chat/internal/pkg/app"
	"go-chat/internal/pkg/app/errcode"
	"go.uber.org/zap"
)

type account struct {
}

func newAccount() *account {
	return &account{}
}

// GetAccountByID
// @Tags     account
// @Summary  获取账户信息
// @accept   application/json
// @Produce  application/json
// @Param    Authorization  header    string               false "x-token 用户令牌"
// @Param    data           query     request.GetAccountByID                   true  "账号信息"
// @Success  200            {object}  common.State{data=reply.GetAccountByID}  "1001:参数有误 1003:系统错误 2009:权限不足 2007:身份不存在 2008:身份验证失败 2010:账号不存在"
// @Router   /api/account/info [get]
func GetAccountByID(c *gin.Context) {
	res := app.NewResponse(c)
	params := &request.GetAccountByID{}
	if err := c.ShouldBindJSON(params); err != nil {
		zap.S().Info("&request.GetAccountByID{} c.ShouldBindJSON(params) failed: %v", err)
		res.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	result, err := logic.Group.Account.GetAccountByID(c, params.AccountID)
	if err != nil {
		res.Reply(err)
	}
	res.Reply(nil, result)
}

func GetAccountListByUserID(c *gin.Context) {

}
