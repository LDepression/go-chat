/**
 * @Author: lenovo
 * @Description:
 * @File:  chat
 * @Version: 1.0.0
 * @Date: 2023/04/30 22:27
 */

package manager

import (
	socketio "github.com/googollee/go-socket.io"
	"sync"
)

func NewChatMap() *ChatMap {
	return &ChatMap{
		m: sync.Map{},
	}
}

type ChatMap struct {
	m   sync.Map // k: accountID v: ConnMap
	sID sync.Map // k: sID v: accountID
}

type ConnMap struct {
	m sync.Map //k: sID v:socketio.Conn
}

// Link 添加设备
func (c *ChatMap) Link(s socketio.Conn, accountID int64) {
	c.sID.Store(s.ID(), accountID) //存入sID和accountID对应关系
	cm, ok := c.m.Load(accountID)
	if !ok {
		cm := &ConnMap{}
		cm.m.Store(s.ID(), s)
		c.m.Store(accountID, cm)
		return
	}
	cm.(*ConnMap).m.Store(s.ID(), s) //这里是另外又去存一个设备
}

// Leave 去除设备
func (c *ChatMap) Leave(s socketio.Conn) {
	accountID, ok := c.sID.LoadAndDelete(s.ID())
	if !ok {
		return
	}
	cm, ok := c.m.Load(accountID)
	if !ok {
		return
	}
	cm.(*ConnMap).m.Delete(s.ID())
	length := 0
	cm.(*ConnMap).m.Range(func(key, value any) bool {
		length++
		return true
	})
	if length == 0 {
		c.m.Delete(accountID)
	}
}

// Send 给指定账号的全部设备推送消息
func (c *ChatMap) Send(accountID int64, event string, args ...interface{}) {
	cm, ok := c.m.Load(accountID)
	if !ok {
		return
	}
	cm.(*ConnMap).m.Range(func(key, value any) bool {
		value.(socketio.Conn).Emit(event, args...) //emit是向event发送消息
		return true
	})
}

// SendMany 给指定多个账号的全部设备推送消息
func (c *ChatMap) SendMany(accountIDs []int64, event string, args ...interface{}) {
	for _, accountID := range accountIDs {
		cm, ok := c.m.Load(accountID)
		if !ok {
			return
		}
		cm.(*ConnMap).m.Range(func(key, value any) bool {
			value.(socketio.Conn).Emit(event, args...)
			return true
		})
	}
}

//HasSID 判断SID是否已经存在

func (c *ChatMap) HasSID(sID string) bool {
	_, ok := c.sID.Load(sID)
	return ok
}
