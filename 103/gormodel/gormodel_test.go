package gormodel

import (
	"fmt"
	inf "gormtest/infrastructure"
	"testing"
)

func TestGormodelCreate(t *testing.T) {
	inf.GormInit()
	GormodelInit()
	GormodelCreate()
}
func TestGormodelFind(t *testing.T) {
	inf.GormInit()
	GormodelInit()
	GormodelFind()
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
