package module

import "fmt"

var ModuleVar int64 = ModuleFun()

func init() {
	fmt.Println("module-- init1 in module.go ")
}
func init() {
	fmt.Println("module-- init2 in module.go ")
}

func ModuleFun() int64 {
	fmt.Println("module-- calling ModuleFun()")
	return 2
}
