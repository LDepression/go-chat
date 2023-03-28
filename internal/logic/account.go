package logic

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/model/reply"
	"go-chat/internal/pkg/app/errcode"
	"time"
)

type account struct {
}

func (account) GetAccountByID(c *gin.Context, accountID int64) (*reply.GetAccountByID, errcode.Err) {

	return &reply.GetAccountByID{
		ID:        0,
		CreatedAt: time.Time{},
		UserID:    0,
		Name:      "",
		Signature: "",
		Avatar:    "",
		Gender:    0,
	}, nil
}
