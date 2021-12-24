package main

import (
	"fmt"

	"gorm.io/gorm"
)

// 用户模型
type User struct {
	ID     uint   // gorm default id
	Name   string // gorm default name
	Age    int    // gorm default age
	Gender string // gorm default gender

	UserDAO string // gorm default user_dao
	DAOUser string // gorm default dao_user

	UserAAA string //gorm default user_aaa
	AAAUser string //gorm default aaa_user

	UserAAIPA string `gorm:"column:user_aaipa"` //gorm default user_aa_ip_a, change to user_aaipa through column
	AAIPAUser string //gorm default aa_ip_a_user
}
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func BasicOperation() {
	// 迁移 schema
	err := gormDB.AutoMigrate(&Product{})
	fmt.Printf("AutoMigrate err:%v\n", err)

	// Create
	dbb := gormDB.Create(&Product{Code: "D42", Price: 100})
	fmt.Printf("Create:%v\n", dbb)

	// Read
	var product Product
	gormDB.First(&product, 1)                 // 根据整形主键查找
	gormDB.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	gormDB.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	gormDB.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	gormDB.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)
}

// 列名约定
func ConventionsColumnName() {
	err := gormDB.AutoMigrate(&User{})
	if nil != err {
		fmt.Printf("AutoMigrate err:%v\n", err)

	}
}

// 列名约定-读取
func ConventionsColumnNameRetrieve() {
	var u User
	gormDB.Debug().Table("users").First(&u)
	fmt.Printf("%+v\n", u)
}
