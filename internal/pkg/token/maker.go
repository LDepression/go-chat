package token

import (
	"errors"
	"time"
)

var ErrSecretLen = errors.New("密钥长度不正确")

type Maker interface {
	// CreateToken 生成Token
	CreateToken(content []byte, expireDate time.Duration) (string, *Payload, error)
	// VerifyToken 解析Token
	VerifyToken(token string) (*Payload, error)
}
