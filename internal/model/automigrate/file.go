/**
 * @Author: lenovo
 * @Description:
 * @File:  file
 * @Version: 1.0.0
 * @Date: 2023/04/12 21:09
 */

package automigrate

import "gorm.io/gorm"

type FType string

var Picture FType = "img"
var Others FType = "file"

type File struct {
	gorm.Model
	URL        string `gorm:" type:varchar(255);not null"`
	FileName   string `gorm:" type:varchar(255);not null"`
	FileType   FType  `gorm:" type:varchar(255);not null;comment:'img,file' " `
	FileSize   uint64 `gorm:"type:bigint;"`
	AccountID  int64
	Account    Account `gorm:"foreignKey:AccountID;references:ID"`
	RelationID *int64
	Relation   Relation `gorm:"foreignKey:RelationID;references:ID"`
}
