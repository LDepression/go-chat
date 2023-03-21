/**
 * @Author: lenovo
 * @Description:
 * @File:  email
 * @Version: 1.0.0
 * @Date: 2023/03/20 20:48
 */

package query

import "context"

var EmailKey = "Email"

func (q *Queries) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	exist, err := q.rdb.SIsMember(ctx, EmailKey, email).Result()
	return exist, err
}
