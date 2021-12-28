package gormodel

import (
	"encoding/json"
	"fmt"
	inf "gormtest/infrastructure"
)

type Gormodel struct {
	ID     int
	Config interface{} `gorm:"type:bytes"` // 将结构体转换为bytes
}

type Student struct {
	Name   string
	Age    int
	Gender string
}

// 初始化
func GormodelInit() {
	inf.GormDB.AutoMigrate(&Gormodel{})
}

// 创建
func GormodelCreate() {
	var stu Student
	stu.Name = "francis"
	stu.Age = 18
	stu.Gender = "male"
	js, e := json.Marshal(stu)
	if nil != e {
		fmt.Printf("%v\n", e)
		return
	}
	model := Gormodel{Config: js}

	inf.GormDB.Debug().Create(&model)
}

// find
func GormodelFind() {
	var model Gormodel
	inf.GormDB.Debug().Find(&model)
	//fmt.Printf("%+v\n", model.Config.(type))

	// 未找到字段为interface，gorm转换的参考文档
	switch t := model.Config.(type) {
	default:
		fmt.Printf("type is %T\n", t)
	}
	a := model.Config.(**interface{})
	fmt.Printf("%+v\n", **a)
	b := **a
	fmt.Printf("b is %+v\n", b.([]byte))

	var stu Student
	e := json.Unmarshal(b.([]byte), &stu)
	if nil != e {
		fmt.Printf("%v\n", e)
		return
	}
	fmt.Printf("stu is %+v\n", stu)
}
