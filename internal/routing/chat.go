/**
 * @Author: lenovo
 * @Description:
 * @File:  chat
 * @Version: 1.0.0
 * @Date: 2023/04/30 22:07
 */

package routing

import (
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	v1 "go-chat/internal/api/v1"
	"go-chat/internal/global"
	chat2 "go-chat/internal/model/chat"
	"time"
)

type ws struct {
}

func (ws) Init(router *gin.Engine) *socketio.Server {
	server := socketio.NewServer(nil)
	{
		server.OnConnect("/", v1.Group.Chat.Handle.OnConnect)
		server.OnError("/", v1.Group.Chat.Handle.OnError)
		server.OnDisconnect("/", v1.Group.Chat.Handle.OnDisconnect)
	}
	go checkTimeOutClients()
	chatHandle(server)
	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))
	return server
}
func chatHandle(server *socketio.Server) {
	event := "/chat"
	server.OnEvent(event, chat2.ClientSendMsg, v1.Group.Chat.Message.SendMsg)
	//server.OnEvent(event, chat2.ClientReadMsg, v1.Group.Chat.Message.ReadMsg)
	server.OnEvent(event, chat2.ClientTest, v1.Group.Chat.Handle.Test)
	server.OnEvent(event, chat2.ClientAuth, v1.Group.Chat.Handle.Auth)
}

func checkTimeOutClients() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			global.ChatMap.CheckForEachAllMap()
		}
	}
}
