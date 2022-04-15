package main

import (
	"fmt"
	"testing"
)

func TestNamedreturn(t *testing.T) {
	a, b := Namedreturn()
	fmt.Printf("a:%+v\n", a)
	fmt.Printf("b:%+v\n", b)
}
