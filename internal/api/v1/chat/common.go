/**
 * @Author: lenovo
 * @Description:
 * @File:  common
 * @Version: 1.0.0
 * @Date: 2023/05/01 0:23
 */

package chat

import (
	socketio "github.com/googollee/go-socket.io"
	"go-chat/internal/dao/mysql/query"
	"go-chat/internal/middleware"
	"go-chat/internal/model"
	"go-chat/internal/model/chat"
	"go-chat/internal/model/common"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app/errcode"
	"go.uber.org/zap"
	"time"
)

// MustAccount 解析token并判断是否是账户，返回token
// 参数: accessToken
// 成功: 解析token的content并进行校验返回*model.Token,nil
// 失败: 返回 myerr.AuthenticationFailed,myerr.UserNotFound,errcode.ErrServer
func MustAccount(accessToken string) (*model.Token, errcode.Err) {
	payload, _, merr := middleware.ParseHeader(accessToken)
	if merr != nil {
		return nil, merr
	}
	content := &model.Content{}
	if err := content.UnMarshal(payload.Content); err != nil {
		return nil, myerr.AuthFailed
	}
	if content.Type != model.AccountToken {
		return nil, myerr.AuthFailed
	}
	qAccount := query.NewQueryAccount()
	accoutInfo, err := qAccount.GetAccountByID(content.ID)
	if err != nil {
		zap.S().Info("err:", err)
		return nil, errcode.ErrServer
	}
	if accoutInfo.ID == 0 {
		return nil, myerr.UserNotFound
	}
	return &model.Token{
		AccessToken: accessToken,
		Payload:     payload,
		Content:     content,
	}, nil
}

// CheckConnCtxToken 检查连接上下文中的token是否有效，有效返回token
// 参数: 连接上下文
// 成功: 上下文中包含 *model.Token 且有效
// 失败: 返回 myerr.AuthenticationFailed,myerr.AuthOverTime
func CheckConnCtxToken(v interface{}) (*model.Token, errcode.Err) {
	token, ok := v.(*model.Token)
	if !ok {
		return nil, myerr.AuthFailed
	}
	if token.Payload.ExpiredAt.Before(time.Now()) {
		return nil, errcode.ErrUnauthorizedTokenTimeout
	}
	return token, nil
}

// CheckAuth 检查token是否有效，有效返回token，否则断开链接
func CheckAuth(s socketio.Conn) (*model.Token, bool) {
	token, err := CheckConnCtxToken(s.Context())
	if err != nil {
		s.Emit(chat.ServerError, common.NewState(err))
		_ = s.Close()
		return nil, false
	}
	return token, true
}
