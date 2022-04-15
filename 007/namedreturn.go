package main

import (
	"errors"
	"fmt"
)

func Namedreturn() (a int, err error) {

	a, errs := 1, errors.New("namedreturn")
	fmt.Printf("errs is %v\n", errs)
	return
}
