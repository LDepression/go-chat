/**
 * @Author: lenovo
 * @Description:
 * @File:  email
 * @Version: 1.0.0
 * @Date: 2023/03/20 23:03
 */

package email

import (
	"errors"
	"fmt"
	"github.com/0RAJA/Rutils/pkg/utils"
	"go-chat/internal/global"
	"go-chat/internal/pkg/email"
	"net/http"
	"sync"
	"time"
)

type Mask struct {
	UserMask sync.Map
	CodeMask sync.Map
	Config   *Config
}

type Config struct {
	SMTPINFO    *email.Email
	DelUserTime time.Duration
	DelCodeTime time.Duration
}

var gMask = Mask{
	UserMask: sync.Map{},
	CodeMask: sync.Map{},
	Config: &Config{
		DelUserTime: global.Settings.Rule.DelUserTime,
		DelCodeTime: global.Settings.Rule.DelCodeTime,
	},
}

type Result struct {
	Code  int
	Error error
}
type SendEmailCode struct {
	Email  string
	Result chan Result
}

func NewSendCodeTask(email string) *SendEmailCode {
	return &SendEmailCode{
		Email:  email,
		Result: make(chan Result, 1),
	}
}

var ErrSendTooMany = errors.New("发送过于频繁")

// CheckEmailBeMask 判断邮箱是否在规定的时间内发送过消息了
func CheckEmailBeMask(email string) bool {
	_, ok := gMask.UserMask.Load(email)
	return ok
}

func (m *Mask) delMark(email string) {
	time.AfterFunc(global.Settings.Rule.DelUserTime, func() { m.UserMask.Delete(email) })
	time.AfterFunc(global.Settings.Rule.DelCodeTime, func() { m.CodeMask.Delete(email) })

}

func (s *SendEmailCode) SendTask() func() {
	return func() {
		//先去判断一下是否已经发送了
		if ok := CheckEmailBeMask(s.Email); ok {
			s.Result <- Result{
				Code:  300,
				Error: ErrSendTooMany,
			}
			return
		}
		//先将用户标记下来
		gMask.UserMask.Store(s.Email, struct{}{})
		sendNewMsg := email.NewEmail(&email.SMTPInfo{
			Host:     global.Settings.SMTPInfo.Host,
			Port:     global.Settings.SMTPInfo.Port,
			IsSSL:    global.Settings.SMTPInfo.IsSSL,
			UserName: global.Settings.SMTPInfo.UserName,
			Password: global.Settings.SMTPInfo.Password,
			From:     global.Settings.SMTPInfo.From,
		})
		code := utils.RandomString(6)
		if err := sendNewMsg.SendMail([]string{s.Email}, fmt.Sprintf("验证码:%s", code), `😘`); err != nil {
			gMask.UserMask.Delete(s.Email)
			s.Result <- Result{
				Code:  http.StatusBadRequest,
				Error: err,
			}
			return
		}
		gMask.CodeMask.Store(s.Email, code)

		//延时删除
		gMask.delMark(s.Email)
		s.Result <- Result{
			Code:  0,
			Error: nil,
		}
		close(s.Result)
		return
	}
}

func (s *SendEmailCode) GetChanResult() Result {
	return <-s.Result
}
