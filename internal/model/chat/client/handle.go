/**
 * @Author: lenovo
 * @Description:
 * @File:  handle
 * @Version: 1.0.0
 * @Date: 2023/05/01 0:45
 */

package client

type TestParams struct {
	Name string `json:"name" validate:"required,gte=1,lte=50"` // 姓名
	Age  string `json:"age" validate:"required,gte=1"`         // 年龄
}

type TestRly struct {
	Name    string `json:"name"`    // 姓名
	Age     string `json:"age"`     // 年龄
	ID      string `json:"id"`      // ID
	Address string `json:"address"` // 地址
}
