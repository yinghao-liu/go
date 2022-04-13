package main

import (
	"errors"
	"fmt"
)

func main() {
	a := errors.New("sss")
	fmt.Printf("%v\n", a)
}
