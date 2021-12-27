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
