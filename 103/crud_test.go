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
