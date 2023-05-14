/**
 * @Author: lenovo
 * @Description:
 * @File:  relation
 * @Version: 1.0.0
 * @Date: 2023/03/30 23:06
 */

package query

import (
	"context"
	"go-chat/internal/pkg/utils"
)

var RelationKey = "Group"

func (q *Queries) AddRelationAccount(ctx context.Context, relationID int64, accountIDs ...int64) error {
	for i := 0; i < len(accountIDs); i++ {
		id := utils.IDToSting(uint(relationID))
		if err := q.rdb.SAdd(ctx, utils.LinkStr(RelationKey, id), accountIDs[i]).Err(); err != nil {
			return err
		}
	}
	return nil
}

// DeleteRelationAccount 从一个群删除多个人
func (q *Queries) DeleteRelationAccount(ctx context.Context, relationID int64, accountIDs ...int64) error {
	if len(accountIDs) == 0 {
		return nil
	}
	id := utils.IDToSting(uint(relationID))
	key := utils.LinkStr(RelationKey, id)
	var ids []interface{}
	for _, id := range accountIDs {
		ids = append(ids, id)
	}
	if err := q.rdb.SRem(ctx, key, ids).Err(); err != nil {
		return err
	}
	return nil
}

func (q *Queries) DeleteAccountByRelations(ctx context.Context, accountID int64, relationIDs []int64) error {
	if len(relationIDs) == 0 {
		return nil
	}
	for _, relationID := range relationIDs {
		id := utils.IDToSting(uint(relationID))
		key := utils.LinkStr(RelationKey, id)
		if err := q.rdb.SRem(ctx, key, accountID).Err(); err != nil {
			return err
		}
	}
	return nil
}

func (q *Queries) GetAllAccountsByRelationID(ctx context.Context, relationID int64) ([]int64, error) {
	id := utils.IDToSting(uint(relationID))
	key := utils.LinkStr(RelationKey, id)
	accountIDStr, err := q.rdb.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var accountIDs []int64
	for _, str := range accountIDStr {
		accountID := utils.StringToIDMust(str)
		accountIDs = append(accountIDs, accountID)
	}
	return accountIDs, nil
}
