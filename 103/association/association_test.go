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
