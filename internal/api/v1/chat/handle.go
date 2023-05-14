/**
 * @Author: lenovo
 * @Description:
 * @File:  handle
 * @Version: 1.0.0
 * @Date: 2023/04/30 22:18
 */

package chat

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"go-chat/internal/global"
	"go-chat/internal/model/chat/client"
	"go-chat/internal/model/common"
	"go-chat/internal/pkg/app/errcode"
	"go-chat/internal/task"
	"go.uber.org/zap"
	"log"
	"time"
)

type handle struct{}

const AuthLimitTimeout = 10 * time.Second

// OnConnect 客户端连接时触发
func (handle) OnConnect(s socketio.Conn) error {
	log.Println("connected:", s.RemoteAddr().String(), s.ID())
	//一段时间后，需要进行AUTH认证，否则就会断开连接
	//time.AfterFunc(AuthLimitTimeout, func() {
	//	if !global.ChatMap.HasSID(s.ID()) {
	//		zap.S().Info(fmt.Sprintf("auth failed:", s.RemoteAddr().String(), s.ID()))
	//		_ = s.Close()
	//	}
	//})
	return nil
}

// OnError 发生错误时触发
func (handle) OnError(s socketio.Conn, e error) {
	log.Println("conn error", e)
	global.ChatMap.Leave(s)
}

func (handle) OnDisconnect(s socketio.Conn, reason string) {
	log.Println(reason)
	global.ChatMap.Leave(s)
}

// Test 测试
func (handle) Test(s socketio.Conn, msg string) string {
	//_, ok := CheckAuth(s)
	//if !ok {
	//	return ""
	//}
	s.SetContext(msg)
	fmt.Println(msg)

	params := &client.TestParams{}

	//log.Println(msg)
	if err := common.Decode(msg, params); err != nil {
		return common.NewState(errcode.ErrParamsNotValid.WithDetails(err.Error())).JsonStr()
	}
	result := common.NewState(nil, client.TestRly{
		Name:    params.Name,
		Age:     params.Age,
		Address: s.RemoteAddr().String(),
		ID:      s.ID(),
	}).JsonStr()
	// test
	s.Emit("test", result)
	return result
}

// Auth 身份验证
func (handle) Auth(s socketio.Conn, accessToken string) string {
	token, merr := MustAccount(accessToken)
	if merr != nil {
		return common.NewState(merr).JsonStr()
	}
	zap.S().Infof(accessToken)
	s.SetContext(token)
	// 加入在线群组
	global.ChatMap.Link(s, int64(token.Content.ID))
	// 通知其他设备
	global.Worker.SendTask(task.AccountLogin(token.AccessToken, s.RemoteAddr().String(), int64(token.Content.ID)))
	log.Println("auth accept:", s.RemoteAddr().String())
	return common.NewState(nil).JsonStr()
}
