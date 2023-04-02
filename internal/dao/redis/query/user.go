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
	"go-chat/internal/global"
	"go-chat/internal/pkg/utils"
)

var UserKey = "user"

func (q *Queries) SaveUserToken(ctx *gin.Context, userID uint, tokens []string) error {
	key := utils.LinkStr(UserKey, utils.IDToSting(userID))
	for _, token := range tokens {
		if err := q.rdb.SAdd(ctx, key, token).Err(); err != nil {
			return err
		}
		q.rdb.Expire(ctx, key, global.Settings.Token.AccessTokenExpire)
	}
	return nil
}

func (q *Queries) CheckUserTokenValid(ctx *gin.Context, userID uint, token string) bool {
	key := utils.LinkStr(UserKey, utils.IDToSting(userID))
	ok := q.rdb.SIsMember(ctx, key, token).Val()
	return ok
}

func (q *Queries) DeleteAllTokenByUser(ctx *gin.Context, userID uint) error {
	key := utils.LinkStr(UserKey, utils.IDToSting(userID))
	if err := q.rdb.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}

func (q *Queries) CountUserToken(ctx *gin.Context, userID uint) int64 {
	key := utils.LinkStr(UserKey, utils.IDToSting(userID))
	count := q.rdb.SCard(ctx, key).Val()
	return count

}
