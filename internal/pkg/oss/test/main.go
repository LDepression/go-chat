/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/04/12 22:37
 */

package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func test() {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := "http://oss-cn-hangzhou.aliyuncs.com"
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyId := "***"
	accessKeySecret := "***"
	bucketName := "lycmall2"
	// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	objectName := "chat/first.jpg"
	// <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
	localFileName := `D:\360download\325021.jpg`
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println(" oss.New failed to create ,err", err)
		panic(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println(" client.Bucket failed to create ,err", err)
		panic(err)
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		fmt.Println(" PutObjectFromFile failed to create ,err", err)
		panic(err)
	}
}

func test2() {
	endpoint := "http://oss-cn-hangzhou.aliyuncs.com"
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyId := "***"
	accessKeySecret := "***"
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("New OSS error:", err)
		panic(err)
	}
	bucket, err := client.Bucket("lycmall2")

	selReq := oss.SelectRequest{}
	// 使用SELECT语句查询文件中的数据。
	selReq.Expression = "select * from ossobject"
	body, err := bucket.SelectObject("chat/first.jpg", selReq)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 读取内容。
	var fc []byte
	_, err = body.Read(fc)
	if err != nil {
		fmt.Println("read error:", err)
	}
	defer body.Close()
	fmt.Println(string(fc))

}

func main() {
	test2()
}
