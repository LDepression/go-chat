/**
 * @Author: lenovo
 * @Description:
 * @File:  message
 * @Version: 1.0.0
 * @Date: 2023/04/30 22:18
 */

package chat

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"go-chat/internal/logic/chat"
	"go-chat/internal/model/chat/client"
	"go-chat/internal/model/chat/serve"
	"go-chat/internal/model/common"
	"go-chat/internal/pkg/app/errcode"
)

type message struct{}

// SendMsg 发送消息
func (message) SendMsg(s socketio.Conn, msg string) string {
	//先来判断一下
	token, ok := CheckAuth(s)
	if !ok {
		return ""
	}
	param := client.HandleSendMsgParams{}
	if err := common.Decode(msg, &param); err != nil {
		return common.NewState(errcode.ErrParamsNotValid.WithDetails(err.Error())).JsonStr()
	}
	//
	fmt.Println(param.RelationID, param.MsgContent)
	result, err := chat.Group.Message.SendMsg(serve.HandleSendMsg{
		AccountID:   int64(token.Content.ID),
		AccessToken: token.AccessToken,
		RelationID:  param.RelationID,
		MsgContent:  param.MsgContent,
		MsgExtend:   param.MsgExtend,
		RlyMsgID:    param.RlyMsgID,
	})
	return common.NewState(err, result).JsonStr()
}

func ReadMsg(s socketio.Conn, reader []int64) {

}
