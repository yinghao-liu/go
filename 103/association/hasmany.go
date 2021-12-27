package ass

import (
	"fmt"
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
	user := HasManyUser{CreditCards: []HasManyCreditCard{
		{Number: "11"},
		{Number: "22"},
	}}

	inf.GormDB.Debug().Create(&user)
}

// 查找
func HasManyFind() {
	var user HasManyUser
	var creditCard []HasManyCreditCard

	user.ID = 1
	// 关联模式
	err := inf.GormDB.Debug().Model(&user).Association("CreditCards").Find(&creditCard)
	if nil != err {
		fmt.Printf("err is %s\n", err.Error())
		return
	}
	fmt.Printf("user is %+v\n", user)
	for i, j := range creditCard {
		fmt.Printf("%d: %+v\n", i, j)

	}

	// 带条件的查找
	var creditCardL []HasManyCreditCard
	err = inf.GormDB.Debug().Model(&user).Where("number like ?", "1_").Association("CreditCards").Find(&creditCardL)
	if nil != err {
		fmt.Printf("err is %s\n", err.Error())
		return
	}
	for i, j := range creditCardL {
		fmt.Printf("%d: %+v\n", i, j)

	}
}
