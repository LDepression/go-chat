package automigrate

type Application struct {
	BaseModel
	AccountID1 uint64  `gorm:"type:bigint;not null"`
	AccountID2 uint64  `gorm:"type:bigint;not null"`
	ApplyMsg   string  `gorm:"type:varchar(50);not null"`
	RefuseMsg  string  `gorm:"type:varchar(50);not null"`
	Status     string  `gorm:"type:varchar(50);comment:ACCEPTED通过,WAITING等待中,REFUSED拒绝"`
	Account1   Account `gorm:"foreignKey:AccountID1;references:ID"`
	Account2   Account `gorm:"foreignKey:AccountID2;references:ID"`
}
