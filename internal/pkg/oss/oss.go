/**
 * @Author: lenovo
 * @Description:
 * @File:  oss
 * @Version: 1.0.0
 * @Date: 2023/04/12 22:25
 */

package oss

import "mime/multipart"

type OssServer interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}
