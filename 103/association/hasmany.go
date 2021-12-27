package ass

import (
	"fmt"
	inf "gormtest/infrastructure"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

// 更新
func HasManyUpdate() {
	var user HasManyUser
	user.ID = 1

	var creditCard HasManyCreditCard
	creditCard.Number = "33"
	// 追加
	err := inf.GormDB.Debug().Model(&user).Association("CreditCards").Append(&creditCard)
	if nil != err {
		fmt.Printf("err is %s\n", err.Error())
		return
	}

	// 替换
	err = inf.GormDB.Debug().Model(&user).Association("CreditCards").Replace(&creditCard)
	if nil != err {
		fmt.Printf("err is %s\n", err.Error())
		return
	}

}

// 删除
func HasManyDelete() {
	var user HasManyUser
	user.ID = 1

	var creditCard HasManyCreditCard
	creditCard.ID = 4
	// 删除
	err := inf.GormDB.Debug().Model(&user).Association("CreditCards").Delete(&creditCard)
	if nil != err {
		fmt.Printf("err is %s\n", err.Error())
		return
	}
	// 清空
	err = inf.GormDB.Debug().Model(&user).Association("CreditCards").Clear()
	if nil != err {
		fmt.Printf("err is %s\n", err.Error())
		return
	}
}

// 关联删除
func HasManyDeleteAss() {
	var user HasManyUser
	user.ID = 1

	var creditCard HasManyCreditCard
	creditCard.ID = 4
	//  删除 user 时，也删除 user 的 CreditCards
	tx := inf.GormDB.Debug().Select("CreditCards").Delete(&user)
	if nil != tx.Error {
		fmt.Printf("err is %s\n", tx.Error.Error())
		return
	}

	// 删除 user 时，也删除用户所有关联记录
	tx = inf.GormDB.Debug().Select(clause.Associations).Delete(&user)
	if nil != tx.Error {
		fmt.Printf("err is %s\n", tx.Error.Error())
		return
	}

}
