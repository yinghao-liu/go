package main

import (
	"fmt"
	"testing"
)

func TestBasicOperation(t *testing.T) {
	gormInit()

	fmt.Printf("----------------TestBasicOperation\n")
	BasicOperation()
}
func TestConventionsColumnName(t *testing.T) {
	gormInit()

	fmt.Printf("----------------TestConventionsColumnName\n")
	ConventionsColumnName()
}
func TestConventionsColumnNameRetrieve(t *testing.T) {
	gormInit()

	fmt.Printf("----------------TestConventionsColumnNameRetrieve\n")
	ConventionsColumnNameRetrieve()
}
