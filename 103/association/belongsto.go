package ass

import (
	inf "gormtest/infrastructure"

	"gorm.io/gorm"
)

// `User` 属于 `Company`，`BelongsToCompanyID` 是外键
// 字段BelongsToCompany名称必须是这个
// 默认情况下，外键的名字，使用拥有者的类型名称加上表的主键的字段名字,本例中拥有者是BelongsToCompany，它的主键是ID
// 默认情况下，引用的是拥有者的主键，本例的是BelongsToCompany的ID字段。可以使用references 来修改它
type BelongsToUser struct {
	gorm.Model
	Name               string
	BelongsToCompanyID int
	BelongsToCompany   BelongsToCompany
}

type BelongsToCompany struct {
	ID   int
	Name string
}

// 重写主键-引用
// 可使用foreignKey指定外键，使用references指定引用
// type User struct {
// 	gorm.Model
// 	Name         string
// 	CompanyRefer string
// 	Company      Company `gorm:"foreignKey:CompanyRefer,references:Code""`
// 	// 使用 CompanyRefer 作为外键
//   }

//   type Company struct {
// 	ID   int
//  Code string
// 	Name string
//   }

// 初始化
func BelongsToInit() {
	inf.GormDB.AutoMigrate(&BelongsToCompany{})
	inf.GormDB.AutoMigrate(&BelongsToUser{})
}

// 创建
func BelongsToCreate() {
	user := BelongsToUser{BelongsToCompany: BelongsToCompany{
		Name: "11",
	}}

	inf.GormDB.Debug().Create(&user)
}
