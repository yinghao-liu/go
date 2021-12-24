package main

import (
	"fmt"
	"testing"
)

func TestAssociation(t *testing.T) {
	gormInit()

	fmt.Printf("----------------TestAssociation\n")
	Association()
}
