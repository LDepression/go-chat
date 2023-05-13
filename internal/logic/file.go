/**
 * @Author: lenovo
 * @Description:
 * @File:  file
 * @Version: 1.0.0
 * @Date: 2023/04/12 23:03
 */

package logic

import (
	"go-chat/internal/dao/mysql/query"
	tx2 "go-chat/internal/dao/mysql/tx"
	"go-chat/internal/model/reply"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app/errcode"
	"go-chat/internal/pkg/oss"
	"go.uber.org/zap"
	"mime/multipart"
)

type file struct{}

func (file) UpdateAccountAvatar(AccountID uint64, file *multipart.FileHeader) (string, errcode.Err) {
	qAccount := query.NewQueryAccount()
	var URL string
	accountInfo, err := qAccount.GetAccountByID(uint(AccountID))
	if err != nil {
		return URL, errcode.ErrServer
	}
	if accountInfo.ID == 0 {
		return URL, myerr.AccountNotExist
	}

	ossClient := oss.NewOss()
	URL, _, err = ossClient.UploadFile(file)
	if err != nil {
		return URL, errcode.ErrServer
	}

	//将文件的相关信息保存到对应的数据库中去
	tx := tx2.NewFileTX()
	if err := tx.UploadAccountAvatar(AccountID, URL, file.Filename, file.Size); err != nil {
		return "", errcode.ErrServer
	}
	return URL, nil
}

func (file) FindFileInfoByFileID(fileID int64) (*reply.FileDetails, errcode.Err) {
	qFile := query.NewFile()
	fileDetails, err := qFile.GetFileByFileID(fileID)
	if err != nil {
		zap.S().Info("qFile.GetFileByFileID(fileID)", zap.Any("err", err))
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}

	var rlp = new(reply.FileDetails)
	rlp.FileType = string(fileDetails.FileType)
	rlp.URL = fileDetails.URL
	rlp.FileName = fileDetails.FileName
	rlp.FileSize = int64(fileDetails.FileSize)

	return rlp, nil
}
