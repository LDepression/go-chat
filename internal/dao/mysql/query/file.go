/**
 * @Author: lenovo
 * @Description:
 * @File:  file
 * @Version: 1.0.0
 * @Date: 2023/04/17 20:31
 */

package query

import (
	"errors"
	"go-chat/internal/dao"
	"go-chat/internal/model/automigrate"
)

type file struct{}

func NewFile() *file {
	return &file{}
}
func (file) GetFileByFileID(fileID int64) (automigrate.File, error) {
	var file automigrate.File
	if result := dao.Group.DB.Model(&automigrate.File{}).Where("id = ?", fileID).First(&file); result.RowsAffected == 0 {
		return file, errors.New("该文件不存在")
	}
	return file, nil
}
