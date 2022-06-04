package scope

import "fmt"

type PublicStruct struct {
	PublicMem  string
	privateMem string
}

type privateStruct struct {
	PublicMem  string
	privateMem string
}

func ScopeTest() {
	var public privateStruct
	var private privateStruct

	fmt.Printf("public.PublicMem:%s\n", public.PublicMem)
	fmt.Printf("public.privateMem:%s\n", public.privateMem)

	fmt.Printf("private.PublicMem:%s\n", private.PublicMem)
	fmt.Printf("private.privateMem:%s\n", private.privateMem)
}
