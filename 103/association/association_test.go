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

func TestAssociationHasMany(t *testing.T) {
	inf.GormInit()
	HasManyInit()

	fmt.Printf("----------------TestAssociationHasMany\n")
	HasManyCreate()
}

func TestAssociationHasManyFind(t *testing.T) {
	inf.GormInit()
	HasManyInit()

	fmt.Printf("----------------TestAssociationHasManyFind\n")
	HasManyFind()
}
