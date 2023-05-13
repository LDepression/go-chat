/**
 * @Author: lenovo
 * @Description:
 * @File:  file
 * @Version: 1.0.0
 * @Date: 2023/04/12 22:54
 */

package request

import "mime/multipart"

type UpdateAvatar struct {
	File *multipart.FileHeader `json:"File"`
}

type FindByFileID struct {
	FileID int64 `json:"FileID" binding:"required"`
}
