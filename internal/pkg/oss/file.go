/**
 * @Author: lenovo
 * @Description:
 * @File:  oss
 * @Version: 1.0.0
 * @Date: 2023/04/12 20:42
 */

package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go-chat/internal/global"
	"mime/multipart"
	"path"
	"time"
)

type OssClient struct {
}

func (*OssClient) UploadFile(file *multipart.FileHeader) (string, string, error) {
	client, err := oss.New(global.Settings.AliyunOSS.Endpoint, global.Settings.AliyunOSS.AccessKeyId, global.Settings.AliyunOSS.AccessKeySecret)
	if err != nil {
		return "", "", err
	}
	bucket, err := client.Bucket(global.Settings.AliyunOSS.BucketName)
	if err != nil {
		return "", "", err
	}
	objectName := global.Settings.AliyunOSS.BucketName + "/" + time.Now().Format("2006-01-02-15:04:05.99") + path.Ext(file.Filename)
	f, err := file.Open()
	if err != nil {
		return "", "", err
	}
	err = bucket.PutObject(objectName, f)
	if err != nil {
		return "", "", err
	}
	return global.Settings.AliyunOSS.BasePath + "/" + objectName, objectName, nil
}

func (*OssClient) DeleteFile(key string) error {
	client, err := oss.New(global.Settings.AliyunOSS.Endpoint, global.Settings.AliyunOSS.AccessKeyId, global.Settings.AliyunOSS.AccessKeySecret)
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(global.Settings.AliyunOSS.BucketName)
	if err != nil {
		return err
	}

	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = bucket.DeleteObject(key)
	if err != nil {
		fmt.Println("delete object failed:", err)
		return err
	}
	return nil
}

func NewOss() OssServer {
	return &OssClient{}
}
