package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func gormInit() {
	if nil == gormDB {
		var err error
		gormDB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		fmt.Printf("gorm.Open:%v\n", gormDB)
	}
}

func main() {
	gormInit()
	Association()
}
