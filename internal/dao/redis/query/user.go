/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/03/25 18:42
 */

package query

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/pkg/utils"
)

var UserKey = "user"

func (q *Queries) SaveUserToken(ctx *gin.Context, userID int64, tokens []string) error {
	key := utils.LinkStr(UserKey, utils.IDToSting(userID))
	for _, token := range tokens {
		if err := q.rdb.SAdd(ctx, key, token).Err(); err != nil {
			return err
		}
	}
	return nil
}

func (q *Queries) CheckUserTokenValid(ctx *gin.Context, userID int64, token string) bool {
	key := utils.LinkStr(UserKey, utils.IDToSting(userID))
	ok := q.rdb.SIsMember(ctx, key, token).Val()
	return ok
}

func (q *Queries) DeleteAllTokenByUser(ctx *gin.Context, userID int64) error {
	key := utils.LinkStr(UserKey, utils.IDToSting(userID))
	if err := q.rdb.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}
