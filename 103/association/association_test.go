package ass

import (
	"fmt"
	inf "gormtest/infrastructure"
	"testing"
)

/******************************test belongs to***************************************/
func TestBelongsToCreate(t *testing.T) {
	inf.GormInit()
	BelongsToInit()

	fmt.Printf("----------------TestBelongsToCreate\n")
	BelongsToCreate()
}

/******************************test has one***************************************/
func TestHasOneCreate(t *testing.T) {
	inf.GormInit()
	HasOneInit()

	fmt.Printf("----------------TestHasOneCreate\n")
	HasOneCreate()
}

/******************************test Has many***************************************/
func TestHasManyCreate(t *testing.T) {
	inf.GormInit()
	HasManyInit()

	fmt.Printf("----------------TestHasManyCreate\n")
	HasManyCreate()
}

func TestHasManyFind(t *testing.T) {
	inf.GormInit()
	HasManyInit()

	fmt.Printf("----------------TestHasManyFind\n")
	HasManyFind()
}

func TestHasManyUpdate(t *testing.T) {
	inf.GormInit()
	HasManyInit()

	fmt.Printf("----------------TestHasManyUpdate\n")
	HasManyUpdate()
}

func TestHasManyDelete(t *testing.T) {
	inf.GormInit()
	HasManyInit()

	fmt.Printf("----------------TestHasManyDelete\n")
	HasManyDelete()
}

func TestHasManyDeleteAss(t *testing.T) {
	inf.GormInit()
	HasManyInit()

	fmt.Printf("----------------TestHasManyDeleteAss\n")
	HasManyDeleteAss()
}

/******************************test Many to Many***************************************/
func TestManyToManyCreate(t *testing.T) {
	inf.GormInit()
	ManyToManyInit()

	fmt.Printf("----------------TestBelongsToCreate\n")
	ManyToManyCreate()
}
func TestManyToManyFind(t *testing.T) {
	inf.GormInit()
	ManyToManyInit()

	fmt.Printf("----------------TestManyToManyFind\n")
	ManyToManyFind()
}
