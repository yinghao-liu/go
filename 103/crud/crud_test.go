package crud

import (
	"fmt"
	inf "gormtest/infrastructure"
	"testing"
)

func TestBasicOperation(t *testing.T) {
	inf.GormInit()

	fmt.Printf("----------------TestBasicOperation\n")
	CRUDBasicOperation()
}
func TestCRUDCreate(t *testing.T) {
	inf.GormInit()
	CRUDInit()

	fmt.Printf("----------------TestCRUDCreate\n")
	CRUDCreate()
}
func TestCRUDFind(t *testing.T) {
	inf.GormInit()
	CRUDInit()

	fmt.Printf("----------------TestCRUDFind\n")
	CRUDFind()
}
