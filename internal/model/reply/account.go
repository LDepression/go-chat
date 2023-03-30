package reply

import (
	"time"
)

type Account struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	Signature string    `json:"signature"`
	Avatar    string    `json:"avatar"`
	Gender    int       `json:"gender"`
}

type GetAccountByID struct {
	Account
}

type GetAccountsByName struct {
	Account []*Account
	Total   int64
}
