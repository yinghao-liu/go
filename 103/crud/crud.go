package crud

import (
	"fmt"
	inf "gormtest/infrastructure"

	"gorm.io/gorm/clause"
)

type Product struct {
	ID    uint `gorm:"primarykey"`
	Code  string
	Price uint
}

func CRUDInit() {
	// 迁移 schema
	err := inf.GormDB.AutoMigrate(&Product{})
	if nil != err {
		fmt.Printf("AutoMigrate err:%v\n", err)
	}
}

func CRUDBasicOperation() {
	// Create
	dbb := inf.GormDB.Create(&Product{Code: "D42", Price: 100})
	fmt.Printf("Create:%v\n", dbb)

	// Read
	var product Product
	inf.GormDB.First(&product, 1)                 // 根据整形主键查找
	inf.GormDB.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	inf.GormDB.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	inf.GormDB.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	inf.GormDB.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)
}

func CRUDCreate() {
	/****************************************************************/
	var ins Product
	ins.ID = 1
	ins.Code = "1"
	ins.Price = 10
	tx := inf.GormDB.Debug().Create(&ins)
	if nil != tx.Error {
		fmt.Printf("Create1 error:%s\n", tx.Error.Error())
	}
	tx = inf.GormDB.Debug().Not("EXITS").Create(&ins)
	if nil != tx.Error {
		fmt.Printf("Create2 error:%s\n", tx.Error.Error())
	}
	inss := make([]Product, 2)
	inss[0] = ins
	inss[1].ID = 2
	inss[1].Code = "2"
	inss[1].Price = 20
	fmt.Printf("inss is %+v\n", inss)

	// 在冲突时，什么都不做
	tx = inf.GormDB.Debug().Clauses(clause.OnConflict{DoNothing: true}).Create(inss)
	if nil != tx.Error {
		fmt.Printf("Create3 error:%s\n", tx.Error.Error())
	}

}
