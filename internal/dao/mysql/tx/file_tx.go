/**
 * @Author: lenovo
 * @Description:
 * @File:  file
 * @Version: 1.0.0
 * @Date: 2023/04/12 23:11
 */

package tx

import (
	"go-chat/internal/dao"
	"go-chat/internal/model/automigrate"
)

type fileTX struct{}

func NewFileTX() *fileTX {
	return &fileTX{}
}

func (f *fileTX) UploadAccountAvatar(accountID uint64, avatarURL string, fileName string, fileSize int64) error {
	tx := dao.Group.DB.Begin()
	//先要把account表中的头像换了，然后将file文件的
	if result := tx.Model(&automigrate.Account{}).Where("id = ?", accountID).Update("avatar", avatarURL); result.RowsAffected == 0 {
		return result.Error
	}
	//然后进行file表的更新
	var file automigrate.File
	file.AccountID = int64(accountID)
	file.FileName = fileName
	file.FileSize = uint64(fileSize)
	file.FileType = automigrate.Picture
	file.URL = avatarURL
	file.RelationID = nil
	if result := tx.Model(&automigrate.File{}).Where("account_id =?", accountID).Save(&file); result.RowsAffected == 0 {
		return result.Error
	}
	tx.Commit()
	return nil
}
