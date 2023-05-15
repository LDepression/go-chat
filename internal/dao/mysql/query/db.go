/**
 * @Author: lenovo
 * @Description:
 * @File:  db
 * @Version: 1.0.0
 * @Date: 2023/05/15 19:31
 */

package query

import (
	"database/sql"
	"go-chat/internal/model/automigrate"
)

type CreateMsgParams struct {
	AccountID  sql.NullInt64
	MsgType    string
	RelationID sql.NullInt64
	MsgContent string
	MsgExtend  *automigrate.MsgExtend
	ReplyMsgID sql.NullInt64
	FileID     sql.NullInt64
}
