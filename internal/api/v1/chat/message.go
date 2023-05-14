/**
 * @Author: lenovo
 * @Description:
 * @File:  message
 * @Version: 1.0.0
 * @Date: 2023/04/30 22:18
 */

package chat

import socketio "github.com/googollee/go-socket.io"

type message struct{}

func (message) SendMsg(s socketio.Conn, msg string) {

}
