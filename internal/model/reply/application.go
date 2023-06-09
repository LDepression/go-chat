/**
 * @Author: lenovo
 * @Description:
 * @File:  application
 * @Version: 1.0.0
 * @Date: 2023/04/04 9:11
 */

package reply

type CreateApplicationRep struct {
	ApplicationID uint64
	Status        string
}
type EasyAccount struct {
	AccountID uint64 `gorm:"type:bigint"`
	Name      string `gorm:"type:varchar(255);not null"`
	Avatar    string `gorm:"type:varchar(255);not null"`
}

type GetApplication struct {
	Applicant *EasyAccount
	Receiver  *EasyAccount
	ApplyMsg  string `gorm:"type:varchar(50);not null"`
	RefuseMsg string `gorm:"type:varchar(50);not null"`
	Status    string `gorm:"type:varchar(50);comment:ACCEPTED通过,WAITING等待中,REFUSED拒绝"`
}

type ApplicationsList struct {
	ApplicationList []*GetApplication
	Total           int64
}
