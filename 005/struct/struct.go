package st

import "fmt"

type base struct {
	Name        string
	Description string
	unexported  string
}

func (b *base) Show() {
	fmt.Printf("base show\n")
}

type Base2 struct{}

func (Base2) show() {
	fmt.Printf("base2 show\n")
}

type BaseInfo struct {
	base
	Base2
	Gender string
}

func (BaseInfo) show() {
	fmt.Printf("BaseInfo show\n")
}
