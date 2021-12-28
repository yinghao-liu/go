package ass

import (
	inf "gormtest/infrastructure"

	"gorm.io/gorm"
)

// HasOneUser 有一张 HasOneCreditCard，HasOneUserID 是外键
// 主体是user，以user视角看
type HasOneUser struct {
	gorm.Model
	HasOneCreditCard HasOneCreditCard
}

type HasOneCreditCard struct {
	gorm.Model
	Number       string
	HasOneUserID uint
}

// 初始化
func HasOneInit() {
	inf.GormDB.AutoMigrate(&HasOneCreditCard{})
	inf.GormDB.AutoMigrate(&HasOneUser{})
}

// 创建
func HasOneCreate() {
	user := HasOneUser{HasOneCreditCard: HasOneCreditCard{
		Number: "11",
	}}

	inf.GormDB.Debug().Create(&user)
}
