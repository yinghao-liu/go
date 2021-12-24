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
func TestConventionsColumnName(t *testing.T) {
	inf.GormInit()

	fmt.Printf("----------------TestConventionsColumnName\n")
	ConventionsColumnName()
}
func TestConventionsColumnNameRetrieve(t *testing.T) {
	inf.GormInit()

	fmt.Printf("----------------TestConventionsColumnNameRetrieve\n")
	ConventionsColumnNameRetrieve()
}
