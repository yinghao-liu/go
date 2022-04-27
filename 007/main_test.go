package main

import (
	"fmt"
	"testing"
)

func TestNamedreturn(t *testing.T) {
	a, err := Namedreturn()
	fmt.Printf("end-- a is %d, err is %#v\n", a, err)
}
