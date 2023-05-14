/**
 * @Author: lenovo
 * @Description:
 * @File:  common
 * @Version: 1.0.0
 * @Date: 2023/03/28 16:56
 */

package automigrate

import (
	"database/sql/driver"
	"encoding/json"
)

func (g FriendType) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb

func (g *FriendType) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

func (g GroupType) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (g *GroupType) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}
