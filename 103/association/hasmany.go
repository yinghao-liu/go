package ass

import (
	inf "gormtest/infrastructure"

	"gorm.io/gorm"
)

// User 有多张 CreditCard，默认HasManyUserID 是外键
type HasManyUser struct {
	gorm.Model
	CreditCards []HasManyCreditCard
}

type HasManyCreditCard struct {
	gorm.Model
	Number        string
	HasManyUserID uint
}

// 初始化
func HasManyInit() {
	inf.GormDB.AutoMigrate(&HasManyUser{})
	inf.GormDB.AutoMigrate(&HasManyCreditCard{})
}

// 创建
func HasManyCreate() {
	HasManyInit()
	usr := HasManyUser{CreditCards: []HasManyCreditCard{
		{Number: "11"},
		{Number: "22"},
	}}

	inf.GormDB.Debug().Create(&usr)
}
