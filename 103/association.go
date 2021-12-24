package main

import (
	"fmt"

	"gorm.io/gorm"
)

// Student 拥有并属于多种 language，`student_languages` 是连接表
type Student struct {
	gorm.Model
	Name      string
	Languages []Language `gorm:"many2many:student_languages;"`
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
	if err := gormDB.AutoMigrate(&Student{}); nil != err {
		fmt.Printf("AutoMigrate err:%v\n", err)
		return err
	}

	usr := Student{
		Name: "francis",
		Languages: []Language{
			{Name: "ZH"},
			{Name: "EN"},
		},
	}

	gormDB.Debug().Create(&usr)

	gormDB.Debug().Save(&usr)
	var lans []Language
	var stud Student
	gormDB.Model(&stud).Association("Languages").Find(&lans)
	for i, j := range lans {
		fmt.Printf("i:%d,j:%+v\n", i, j)
	}
	return nil
}
