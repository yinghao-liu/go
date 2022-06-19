package scope_test

import (
	"fmt"
	"function/scope"
	"testing"
)

type Sub struct {
	scope.PublicStruct
}

func TestScope(t *testing.T) {
	var public scope.PublicStruct
	// var private scope.privateStruct // privateStruct not exported by package scope

	fmt.Printf("public.PublicMem:%s\n", public.PublicMem)
	//fmt.Printf("public.privateMem:%s\n", public.privateMem) // public.privateMem undefined

	// fmt.Printf("private.PublicMem:%s\n", private.PublicMem)
	// fmt.Printf("private.privateMem:%s\n", private.privateMem)

	//
	scope.ScopeTest()

	// Sub
	var sub Sub
	fmt.Printf("public.PublicMem:%s\n", sub.PublicMem)
	// fmt.Printf("private.privateMem:%s\n", sub.privateMem) // sub.privateMem undefined
}

func TestScopeTest(t *testing.T) {
	scope.ScopeAndMultiAssign()
	scope.ScopeGlobal()
	fmt.Printf("-----------------------\n")
	scope.ScopeGlobalAssign()
	scope.ScopeGlobal()
}
