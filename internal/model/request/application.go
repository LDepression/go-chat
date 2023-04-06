package request

import "go-chat/internal/model/common"

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
