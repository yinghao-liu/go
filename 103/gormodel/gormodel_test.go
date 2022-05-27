package gormodel

import (
	"fmt"
	inf "gormtest/infrastructure"
	"testing"
)

/***************************************Gormodel with type-byte************************************************/
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
func TestGormodelDelete(t *testing.T) {
	inf.GormInit()
	StudentInit()
	StudentCreate()
	StudentDelete()

}
func TestGormodelCreateV2(t *testing.T) {
	inf.GormInit()
	GormodelInit()
	GormodelCreateV2()
}
func TestGormodelFindV2(t *testing.T) {
	inf.GormInit()
	GormodelInit()
	GormodelFindV2()
}

/***************************************Conventions************************************************/
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

/**************************Compatible*********************************/
func TestGormodelCompatibleCreate(t *testing.T) {
	inf.GormInit()

	fmt.Printf("----------------TestGormodelCompatibleCreate\n")
	GormodelCompatibleCreate()
}
