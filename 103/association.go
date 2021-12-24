package main

import (
	"fmt"

	"gorm.io/gorm"
)

// User 拥有并属于多种 language，`user_languages` 是连接表
type User struct {
	gorm.Model
	Name      string
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

func Association() error {

	// 迁移 schema
	if err := gormDB.AutoMigrate(&Language{}); nil != err {
		fmt.Printf("AutoMigrate err:%v\n", err)
		return err
	}
	if err := gormDB.AutoMigrate(&User{}); nil != err {
		fmt.Printf("AutoMigrate err:%v\n", err)
		return err
	}

	usr := User{
		Name: "francis",
		Languages: []Language{
			{Name: "ZH"},
			{Name: "EN"},
		},
	}

	gormDB.Debug().Create(&usr)

	gormDB.Debug().Save(&usr)
	var lans []Language
	var usr2 User
	gormDB.Model(&usr2).Association("Languages").Find(&lans)
	//db.Table("users").Association("Languages").Find(&lans)
	for i, j := range lans {
		fmt.Printf("i:%d,j:%+v\n", i, j)
	}
	return nil
}
