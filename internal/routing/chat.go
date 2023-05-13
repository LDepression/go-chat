/**
 * @Author: lenovo
 * @Description:
 * @File:  chat
 * @Version: 1.0.0
 * @Date: 2023/04/30 22:07
 */

package routing

import (
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	v1 "go-chat/internal/api/v1"
	"go-chat/internal/middleware"
	chat2 "go-chat/internal/model/chat"
	"log"
	"net/http"
)

type ws struct {
}

func (ws) Init(router *gin.Engine) *socketio.Server {
	server := socketio.NewServer(nil)
	// redis 适配器
	ok, err := server.Adapter(&socketio.RedisAdapterOptions{
		Addr:    "127.0.0.1:6379",
		Prefix:  "socket.io",
		Network: "tcp",
	})

	fmt.Println("redis:", ok)

	if err != nil {
		log.Fatal("error:", err)
	}

	// 连接成功
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		// 申请一个房间
		s.Join("bcast")
		fmt.Println("连接成功：", s.ID())
		return nil
	})

	// 接收”notice“事件
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		s.Emit("reply", "have "+msg) // 回复内容
		log.Println("notice收到内容：:", msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		fmt.Println("=====chat====>", msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn, msg string) string {
		last := s.Context().(string)
		s.Emit("bye", msg)
		fmt.Println("============>", last)
		//s.Close()
		return last
	})

	//----------------------------房间操作----------------------------------
	server.OnEvent("/", "join", func(s socketio.Conn, msg string) {
		//ok= server.JoinRoom("/","bcast",s)
		// fmt.Println(msg,"==房间操作==ok===",ok)
		server.BroadcastToRoom("", "bcast", "event:name", msg)
	})

	router.GET("/room", func(context *gin.Context) {
		// 向房间内的所有人员发消息
		server.BroadcastToRoom("", "bcast", "event:name", "通知")
		fmt.Println("=========向房间内的所有人员发消息======>")
	})

	//----------------------------------end----------------------------------

	// 连接错误
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("连接错误:", e)

	})
	// 关闭连接
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("关闭连接：", reason)
	})
	chatHandle(server)

	router.Use(gin.Recovery(), middleware.Cors())
	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))
	router.StaticFS("/public", http.Dir("../asset"))
	return server
}
func chatHandle(server *socketio.Server) {
	event := "/chat"
	//server.OnEvent(event, chat2.ClientSendMsg, v1.Group.Chat.Message.SendMsg)
	//server.OnEvent(event, chat2.ClientReadMsg, v1.Group.Chat.Message.ReadMsg)
	server.OnEvent(event, chat2.ClientTest, v1.Group.Chat.Handle.Test)
	server.OnEvent(event, chat2.ClientAuth, v1.Group.Chat.Handle.Auth)
}
