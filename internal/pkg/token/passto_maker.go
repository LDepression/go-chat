package token

import (
	"errors"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto *paseto.V2
	key    []byte
}

func NewPasetoMaker(key []byte) (Maker, error) {
	if len(key) != chacha20poly1305.KeySize {
		return nil, ErrSecretLen
	}
	return &PasetoMaker{
		paseto: paseto.NewV2(),
		key:    key,
	}, nil
}

func (p *PasetoMaker) CreateToken(content []byte, expireDate time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(content, expireDate)
	if err != nil {
		return "", nil, nil
	}
	token, err := p.paseto.Encrypt(p.key, payload, nil)
	if err != nil {
		return "", nil, err
	}
	return token, payload, nil
}

func (p *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := p.paseto.Decrypt(token, p.key, payload, nil)
	if err != nil {
		return nil, err
	}
	if payload.ExpiredAt.Before(time.Now()) {
		return nil, errors.New("超时错误")
	}
	return payload, nil
}
