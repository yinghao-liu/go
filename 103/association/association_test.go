package ass

import (
	"fmt"
	inf "gormtest/infrastructure"
	"testing"
)

func TestAssociation(t *testing.T) {
	inf.GormInit()

	fmt.Printf("----------------TestAssociation\n")
	Association()
}

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
func TestAssociationHasManyCreate(t *testing.T) {
	inf.GormInit()
	HasManyInit()

	fmt.Printf("----------------TestAssociationHasManyCreate\n")
	HasManyCreate()
}

func TestAssociationHasManyFind(t *testing.T) {
	inf.GormInit()
	HasManyInit()

	fmt.Printf("----------------TestAssociationHasManyFind\n")
	HasManyFind()
}

func TestAssociationHasManyUpdate(t *testing.T) {
	inf.GormInit()
	HasManyInit()

	fmt.Printf("----------------TestAssociationHasManyUpdate\n")
	HasManyUpdate()
}

func TestAssociationHasManyDelete(t *testing.T) {
	inf.GormInit()
	HasManyInit()

	fmt.Printf("----------------TestAssociationHasManyDelete\n")
	HasManyDelete()
}

func TestAssociationHasManyDeleteAss(t *testing.T) {
	inf.GormInit()
	HasManyInit()

	fmt.Printf("----------------TestAssociationHasManyDeleteAss\n")
	HasManyDeleteAss()
}
