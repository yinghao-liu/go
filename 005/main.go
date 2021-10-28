package main

import "fmt"

type Base struct {
	Name        string
	Description string
}

type BaseInfo struct {
	Base
	Gender string
}

func main() {
	var a BaseInfo
	a.Name = "abc"
	a.Description = "cc"
	a.Gender = "male"
	fmt.Printf("%+v\n", a.Base)
}
