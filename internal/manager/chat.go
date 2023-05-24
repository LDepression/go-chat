/**
 * @Author: lenovo
 * @Description:
 * @File:  chat
 * @Version: 1.0.0
 * @Date: 2023/04/30 22:27
 */

package manager

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"sync"
	"time"
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
	m sync.Map //k: sID v:Active
}

type Active struct {
	s          socketio.Conn
	activeTime time.Time
}

// Link 添加设备
func (c *ChatMap) Link(s socketio.Conn, accountID int64) {
	c.sID.Store(s.ID(), accountID) //存入sID和accountID对应关系
	cm, ok := c.m.Load(accountID)
	if !ok {
		cm := &ConnMap{}
		activeConn := &Active{}
		activeConn.s = s
		activeConn.activeTime = time.Now()
		cm.m.Store(s.ID(), activeConn)
		c.m.Store(accountID, cm)
		return
	}
	activeConn := &Active{}
	activeConn.s = s
	activeConn.activeTime = time.Now()
	cm.(*ConnMap).m.Store(s.ID(), activeConn) //这里是另外又去存一个设备
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
		t := value.(*Active)
		fmt.Println(args)
		t.s.Emit(event, args...) //emit是向event发送消息
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
			t := value.(Active)
			t.s.Emit(event, args...)
			return true
		})
	}
}

//HasSID 判断SID是否已经存在

func (c *ChatMap) HasSID(sID string) bool {
	_, ok := c.sID.Load(sID)
	return ok
}

func (c *ChatMap) CheckForEachAllMap() {

	//fmt.Println("**************************************")

	c.m.Range(func(key, value any) bool {
		//key就是account,value就是ConnMap
		value.(*ConnMap).m.Range(func(key1, value1 any) bool {
			//此时的key1就是sID,value1就是Active
			activeTime := value1.(*Active).activeTime
			if time.Now().Sub(activeTime) > 10*time.Minute {
				err := value1.(*Active).s.Close()
				if err != nil {
					return false
				}
			}
			return true
		},
		)
		return true
	},
	)
}
