package main

import (
	"fmt"
	"testing"
)

func TestNamedreturn(t *testing.T) {
	a, err := Namedreturn()
	fmt.Printf("end-- a is %d, err is %#v\n", a, err)
}

func TestVarArg(t *testing.T) {
	var slc = []int{1, 2, 3}
	//VarArgs("aa", slc)   // cannot use slc (variable of type []int) as int value in argument to VarArgs
	SliceArgs("aa", slc) // OK

	VarArgs("aa", 1, 2, 3) // ok
	//SliceArgs("aa", 1, 2, 3) // too many arguments in call to SliceArgs	have(string, number, number, number) want(string, []int)

	VarArgs("aa", slc...) // OK
}
