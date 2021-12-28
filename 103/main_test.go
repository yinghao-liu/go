package main

import (
	"gormtest/gormodel"
	inf "gormtest/infrastructure"
	"testing"
)

func TestGormodelCreate(t *testing.T) {
	inf.GormInit()
	gormodel.GormodelInit()
	gormodel.GormodelCreate()
}
func TestGormodelFind(t *testing.T) {
	inf.GormInit()
	gormodel.GormodelInit()
	gormodel.GormodelFind()
}
