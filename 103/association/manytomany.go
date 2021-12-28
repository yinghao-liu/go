package ass

import (
	"fmt"
	inf "gormtest/infrastructure"

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

// 初始化
func ManyToManyInit() error {
	// 迁移 schema
	if err := inf.GormDB.AutoMigrate(&Language{}); nil != err {
		fmt.Printf("AutoMigrate err:%v\n", err)
		return err
	}
	if err := inf.GormDB.AutoMigrate(&Student{}); nil != err {
		fmt.Printf("AutoMigrate err:%v\n", err)
		return err
	}
	return nil
}

// 创建
func ManyToManyCreate() {
	usr := Student{
		Name: "francis",
		Languages: []Language{
			{Name: "ZH"},
			{Name: "EN"},
		},
	}

	inf.GormDB.Debug().Create(&usr)
	inf.GormDB.Debug().Save(&usr)
}

// 查找
func ManyToManyFind() error {

	var lans []Language
	var stud Student
	inf.GormDB.Model(&stud).Association("Languages").Find(&lans)
	for i, j := range lans {
		fmt.Printf("i:%d,j:%+v\n", i, j)
	}
	return nil
}
