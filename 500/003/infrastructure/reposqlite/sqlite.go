package reposqlite

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbName = "database.db"

var dbHandle *gorm.DB

func DatabaseInit() error {
	if nil != dbHandle {
		return nil
	}
	dbHandle, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Printf("dbHandle is %p\n", dbHandle)
	return nil
}

func DatabaseHandle() error {
	fmt.Printf("dbHandle is %p\n", dbHandle)
	return nil
}
