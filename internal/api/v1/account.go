/**
 * @Author: lenovo
 * @Description:
 * @File:  account
 * @Version: 1.0.0
 * @Date: 2023/03/28 16:44
 */

package v1

type account struct{}

// CreateAccount 创建账户
// @Tags     email
// @Summary  创建账户
// @accept   application/json
// @Produce  application/json
// @Param 	Authorization 	header 	string 	false 	"x-token 用户令牌"
// @Param   data  body      request.CheckEmailExist  true  "email"
// @Success  200   {object}  common.State{}     "1001:参数有误 1003:系统错误 "
// @Router   /api/v1/email/check [post]
//func (account) CreateAccount(ctx *gin.Context) {
//	rly := app.NewResponse(ctx)
//	var req request.CreateAccountReq
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		base.HandleValidatorError(ctx, err)
//		zap.S().Infof("ctx.ShouldBindJSON(&req) failed: %v", err)
//		return
//	}
//
//	logic.Group.Account.CreateAccount(ctx, req)
//}
