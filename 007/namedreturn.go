package main

import (
	"errors"
	"fmt"
)

func Namedreturn() (a int, err error) {

	// 命名返回在同一作用域下时，:= a是赋值，errs是新建
	a, errs := 1, errors.New("namedreturn")
	fmt.Printf("1-- a is %d, errs is %s\n", a, errs.Error())

	// 命名返回不在同一作用域下时, := 均是新建
	if a, err := 2, errors.New("in if scope"); a > 1 {
		fmt.Printf("2-- a is %d, err is %s\n", a, err.Error())
	}
	return
	// if a, err = 3, errors.New("in 3 if scope"); a > 1 {
	// 	fmt.Printf("3-- a is %d, err is %s\n", a, err.Error())
	// }
	// return
}
