package main

import "fmt"

type Base struct {
	Name        string
	Description string
}

func (b *Base) show() {
	fmt.Printf("base show\n")
}

type Base2 struct{}

func (Base2) show() {
	fmt.Printf("base2 show\n")
}

type BaseInfo struct {
	Base
	Base2
	Gender string
}

func (BaseInfo) show() {
	fmt.Printf("BaseInfo show\n")
}

// alias
type BaseAlias Base

func main() {
	var a BaseInfo
	a.Name = "abc"
	a.Description = "cc"
	a.Gender = "male"
	a.show() // if BaseInfo did not declare show(), there is an error: ambiguous selector a.show
	fmt.Printf("%+v\n", a.Base)

	// alias
	//var b BaseAlias
	//b.show() // b.show undefined (type Base2 has no field or method show)
}
