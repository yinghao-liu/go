package main

import "fmt"

// 变长参数
func VarArgs(a string, args ...int) {
	fmt.Printf("a is %T,args is %T\n", a, args) // a is string,args is []int
}

// 切片参数
func SliceArgs(a string, args []int) {
	fmt.Printf("a is %T,args is %T\n", a, args) // a is string,args is []int
}
