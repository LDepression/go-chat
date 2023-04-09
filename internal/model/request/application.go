/**
 * @Author: lenovo
 * @Description:
 * @File:  application
 * @Version: 1.0.0
 * @Date: 2023/04/04 8:35
 */

package request

type CreateApplicationReq struct {
	AccountID      uint64 `json:"accountID" binding:"required"`
	ApplicationMsg string `json:"applicationMsg" binding:"required"`
}

type DeleteApplicationReq struct {
	AccountID uint64 `json:"accountID" binding:"required"`
}
