package reply

import (
	"time"
)

type AccountInfo struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	Signature string    `json:"signature"`
	Avatar    string    `json:"avatar"`
	Gender    int       `json:"gender"`
}

type GetAccountByID struct {
	AccountInfo
}

type GetAccountsByName struct {
	AccountInfos []*AccountInfo
	Total        int64
}

type GetAccountsByUserID struct {
	AccountInfos []*AccountInfo
	Total        int64
}
