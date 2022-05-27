package gormodel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	inf "gormtest/infrastructure"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

/*************************************model with interface{}*****************************************/
// model with interface{} - type:bytes
type Gormodel struct {
	ID     int
	Config interface{} `gorm:"type:bytes"` // 将结构体转换为json bytes再存储
}

type Student struct {
	gorm.Model
	Name   string
	Age    int
	Gender string
}

// 兼容性测试
type CompatibleStudent struct {
	Name string
	Age  int
	//Gender string
	Gender int // 兼容性测试
}

/*************************************Conventions*****************************************/
// 列名约定
func ConventionsColumnName() {
	err := inf.GormDB.AutoMigrate(&User{})
	if nil != err {
		fmt.Printf("AutoMigrate err:%v\n", err)

	}
}

// 列名约定-读取
func ConventionsColumnNameRetrieve() {
	var u User
	inf.GormDB.Debug().Table("users").First(&u)
	fmt.Printf("%+v\n", u)
}

/*********************************************student********************************************************/
// 初始化
func StudentInit() {
	inf.GormDB.AutoMigrate(&Student{})
}

// 创建
func StudentCreate() {
	var stu Student
	stu.Name = "francis"
	stu.Age = 18
	stu.Gender = "male"

	inf.GormDB.Debug().Create(&stu)
	fmt.Printf("stu is %#v\n", stu)
}

// delete
func StudentDelete() {
	// 返回所有列  sqlite不支持
	var stu []Student
	inf.GormDB.Clauses(clause.Returning{}).Debug().Where("name = ?", "francis").Delete(&stu)
	// DELETE FROM `users` WHERE role = "admin" RETURNING *
	// users => []User{{ID: 1, Name: "jinzhu", Role: "admin", Salary: 100}, {ID: 2, Name: "jinzhu.2", Role: "admin", Salary: 1000}}
	fmt.Printf("stu is %#v\n", stu)
	// 返回指定的列
	//inf.GormDB.Clauses(clause.Returning{Columns: []clause.Column{{Name: "name"}, {Name: "salary"}}}).Where("role = ?", "admin").Debug().Delete(&users)
	// DELETE FROM `users` WHERE role = "admin" RETURNING `name`, `salary`
	// users => []User{{ID: 0, Name: "jinzhu", Role: "", Salary: 100}, {ID: 0, Name: "jinzhu.2", Role: "", Salary: 1000}}

	//inf.GormDB.Debug().Find(&model)
}

/***************************************Gormodel with type-byte************************************************/
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

	c := **model.Config.(**interface{})
	fmt.Printf("c is %+v\n", c.([]byte))

	var stu Student
	e := json.Unmarshal(b.([]byte), &stu)
	if nil != e {
		fmt.Printf("%v\n", e)
		return
	}
	fmt.Printf("stu is %+v\n", stu)
}

/***************************************Gormodel v2************************************************/
func (s *Student) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var result Student
	if err := json.Unmarshal(bytes, &result); nil != err {
		return errors.New(err.Error())
	}

	*s = result
	return nil
}

func (s Student) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// 创建
func GormodelCreateV2() {
	var stu Student
	stu.Name = "francis"
	stu.Age = 18
	stu.Gender = "male"

	model := Gormodel{Config: stu}

	inf.GormDB.Debug().Create(&model)
}

// find
func GormodelFindV2() {
	var model Gormodel
	inf.GormDB.Debug().Find(&model)
	//fmt.Printf("%+v\n", model.Config.(type))

	// 未找到字段为interface，gorm转换的参考文档
	// type is **interface {}
	switch t := model.Config.(type) {
	default:
		fmt.Printf("type is %T\n", t)
	}
}

/*************************************Compatible***************************************/

// 创建
func GormodelCompatibleCreate() {
	inf.GormDB.Debug().AutoMigrate(&CompatibleStudent{})

	var stu CompatibleStudent
	stu.Name = "francis"
	stu.Age = 18
	stu.Gender = 2

	inf.GormDB.Debug().Create(stu)
	var stus []CompatibleStudent
	inf.GormDB.Debug().Find(&stus)
	fmt.Printf("stus:%+v\n", stus)
}
