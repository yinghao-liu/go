package crud

import (
	"fmt"
	inf "gormtest/infrastructure"
	"testing"
)

func TestBasicOperation(t *testing.T) {
	inf.GormInit()

	fmt.Printf("----------------TestBasicOperation\n")
	BasicOperation()
}
