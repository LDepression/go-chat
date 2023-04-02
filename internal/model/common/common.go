/**
 * @Author: lyc
 * @Description:
 * @File:  common
 * @Version: 1.0.0
 * @Date: 2023/03/20 14:56
 */

package common

import "time"

// State 状态码
type State struct {
	Code int         `json:"status_code"`    // 状态码，0-成功，其他值-失败
	Msg  string      `json:"status_msg"`     // 返回状态描述
	Data interface{} `json:"data,omitempty"` // 失败时返回空
}

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
