/**
 * @Author: lenovo
 * @Description:
 * @File:  application
 * @Version: 1.0.0
 * @Date: 2023/04/04 8:35
 */

package request

import "go-chat/internal/model/common"

type CreateApplicationReq struct {
	AccountID      uint64 `json:"accountID" binding:"required"`
	ApplicationMsg string `json:"applicationMsg" binding:"required"`
}

type DeleteApplicationReq struct {
	AccountID uint64 `json:"accountID" binding:"required"`
}

type AcceptApplication struct {
	ApplicantID uint `json:"applicant_id" form:"applicant_id" binding:"required"`
}

type RefuseApplication struct {
	ApplicantID uint   `json:"applicant_id" form:"applicant_id" binding:"required"`
	RefuseMsg   string `json:"refuse_msg" form:"refuse_msg"`
}

type GetApplicationsList struct {
	common.Pager
}
