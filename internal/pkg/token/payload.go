package token

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	// 用于管理每个JWT
	ID      uuid.UUID
	Content []byte //可以是用户或者是账户
	// 创建时间用于检验
	IssuedAt  time.Time `json:"issued-at"`
	ExpiredAt time.Time `json:"expired-at"`
}

func NewPayload(content []byte, expireDate time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &Payload{
		ID:        tokenID,
		Content:   content,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(expireDate),
	}, nil
}
