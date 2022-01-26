package main

import (
	"fmt"
	_ "functionextension/module"
)

var T int64 = a()

func init() {
	fmt.Println("main-- init in main.go ")
}

func a() int64 {
	fmt.Println("main-- calling a()")
	return 2
}

func main() {
	fmt.Println("main-- calling main")
	//module.ModuleFun()
	plunginInit()
}
