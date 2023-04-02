/**
 * @Author: lenovo
 * @Description:
 * @File:  token
 * @Version: 1.0.0
 * @Date: 2023/03/28 17:58
 */

package model

import "encoding/json"

type TokenType string

const (
	UserToken    TokenType = "user"
	AccountToken TokenType = "account"
)

type Content struct {
	Type TokenType `json:"type"`
	ID   uint      `json:"id"`
}

func NewContent(t TokenType, id uint) *Content {
	return &Content{
		Type: t,
		ID:   id,
	}
}
func (c *Content) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Content) UnMarshal(data []byte) error {
	err := json.Unmarshal(data, &c)
	return err
}
