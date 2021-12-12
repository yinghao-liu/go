package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

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

var db *gorm.DB

func Association() error {

	// 迁移 schema
	if err := db.AutoMigrate(&Language{}); nil != err {
		fmt.Printf("AutoMigrate err:%v\n", err)
		return err
	}
	if err := db.AutoMigrate(&User{}); nil != err {
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

	db.Debug().Create(&usr)

	db.Debug().Save(&usr)
	var lans []Language
	var usr2 User
	db.Model(&usr2).Association("Languages").Find(&lans)
	//db.Table("users").Association("Languages").Find(&lans)
	for i, j := range lans {
		fmt.Printf("i:%d,j:%+v\n", i, j)
	}
	return nil
}

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Printf("gorm.Open:%v\n", db)
	Association()

}

func testt() {
	// 迁移 schema
	err := db.AutoMigrate(&Product{})
	fmt.Printf("AutoMigrate err:%v\n", err)

	// Create
	dbb := db.Create(&Product{Code: "D42", Price: 100})
	fmt.Printf("Create:%v\n", dbb)

	// Read
	var product Product
	db.First(&product, 1)                 // 根据整形主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)
}
