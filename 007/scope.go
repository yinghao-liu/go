package main

import "fmt"

var global int

func ScopeAndMultiAssign() {
	global, num := 10, 100
	fmt.Printf("global in local is: %d, num: %d\n", global, num)
}
func ScopeGlobal() {
	fmt.Printf("global in global is: %d\n", global)
}
func ScopeGlobalAssign() {
	global = 99
}
