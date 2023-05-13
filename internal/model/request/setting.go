/**
 * @Author: lenovo
 * @Description:
 * @File:  setting
 * @Version: 1.0.0
 * @Date: 2023/04/18 20:54
 */

package request

type GetPinsReq struct {
	AccountID uint64 `json:"AccountID" binding:"required"`
}

type UpdatePinsReq struct {
	IsPin      bool  `json:"IsPin"`
	RelationID int64 `json:"RelationID" binding:"required"`
}

type UpdateNickName struct {
	NickName   string `json:"NickName" binding:"required"`
	RelationID int64  `json:"RelationID" binding:"required"`
}

type UpdateIsDisturbState struct {
	IsDisturbState bool  `json:"IsDisturbState"`
	RelationID     int64 `json:"RelationID" binding:"required"`
}

type UpdateShowState struct {
	RelationID int64 `json:"RelationID" `
	IsShow     bool  `json:"IsShow" binding:"required"`
}
