/**
 * @Author: lyc
 * @Description:
 * @File:  common
 * @Version: 1.0.0
 * @Date: 2023/03/20 14:56
 */

package common

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"go-chat/internal/pkg/app/errcode"
	"go.uber.org/zap"
	"sync"
	"time"
)

// State 状态码
type State struct {
	Code int         `json:"status_code"`    // 状态码，0-成功，其他值-失败
	Msg  string      `json:"status_msg"`     // 返回状态描述
	Data interface{} `json:"data,omitempty"` // 失败时返回空
}

// NewState 创建一个标准回复格式
// 参数: merr 错误信息(可为nil) datas 数据(存在只选择第一个值)
// 返回: 标准的回复格式结构
func NewState(merr errcode.Err, datas ...interface{}) *State {
	var data interface{}
	if len(datas) > 0 {
		data = datas[0]
	}
	if merr == nil {
		merr = errcode.StatusOK
	} else {
		data = nil
	}
	return &State{
		Code: merr.ECode(),
		Msg:  merr.Error(),
		Data: data,
	}
}

// Json 将结构体转换为json格式的数据
func (s *State) Json() ([]byte, error) {
	return json.Marshal(s)
}

// JsonStr 将结构体转换为json格式的数据，如果出错，则返回空json字符串
func (s *State) JsonStr() string {
	b, err := s.Json()
	if err != nil {
		zap.S().Infof("err:%v", err)
		return "{}"
	}
	return string(b)
}

type ApplicationStatus string
type RelationType string

const (
	RelationTypeGroup  RelationType = "group"
	RelationTypeFriend RelationType = "friend"
	RelationTypeNo     RelationType = ""
)

const (
	ApplicationStateAccepted ApplicationStatus = "已接受"
	ApplicationStateLoading  ApplicationStatus = "申请中"
	ApplicationStateRefused  ApplicationStatus = "已拒绝"
)

type RelationType string

const (
	FriendType RelationType = "friend"
	GroupType  RelationType = "group"
)

type List struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
}

type Pager struct {
	Page     int32 `json:"Page" form:"Page"`
	PageSize int32 `json:"PageSize" form:"PageSize"`
}

type Token struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

var validate *validator.Validate
var validateOnce sync.Once

// Decode 将json格式的数据解析到结构体,并进行校验
// 参数: data: json格式的数据，v: 绑定的结构体
// 返回: 错误信息
func Decode(data string, v interface{}) error {
	if err := json.Unmarshal([]byte(data), v); err != nil {
		return err
	}
	validateOnce.Do(func() { validate = validator.New() })
	return validate.Struct(v)
}
