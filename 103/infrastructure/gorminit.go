package inf

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func GormInit() {
	if nil == GormDB {
		var err error
		GormDB, err = gorm.Open(sqlite.Open("gormtest.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		fmt.Printf("gorm.Open:%v\n", GormDB)
	}
}
