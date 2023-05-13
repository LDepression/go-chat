/**
 * @Author: lenovo
 * @Description:
 * @File:  file
 * @Version: 1.0.0
 * @Date: 2023/04/17 20:28
 */

package reply

type FileDetails struct {
	FileSize int64  `json:"FileSize"`
	FileName string `json:"FileName"`
	URL      string `json:"URL"`
	FileType string `json:"FileType"`
}

type UploadReply struct {
	Avatar string `json:"Avatar"`
}
