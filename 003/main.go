package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Circle struct {
	radius float32
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

type Shaper interface {
	Area() float32
}

func typeTest() {
	var any interface{}

	fmt.Printf("'any' is [%T]-[%p]-[%p]\n", any, &any, any)

	sq1 := new(Square)
	any = sq1
	fmt.Printf("'any' is [%T]-[%p]-[%p]\n", any, &any, any)

	var sq2 Square
	any = sq2
	fmt.Printf("'any' is [%T]-[%p]-[%p]\n", any, &any, any)

	// t and any are not same value, but point to the same address, if ok is true
	// test if any(*Square) suit interface Shaper, if true, return *Square
	if t, ok := any.(Shaper); ok {
		fmt.Printf("dynamic value of 't' with value  %T %v %v\n", t, t, ok)
		fmt.Printf("'t' is %T-%p-%p\n", t, &t, t)
	} else {
		fmt.Printf("dynamic value of 't' without value  %T %v %v\n", t, t, ok)
	}

	/* 	if t, ok := any.(*Square); ok {
	   		fmt.Printf("dynamic value of 't' with value  %T %v %v\n", t, t, ok)
	   		fmt.Printf("'t' is %T-%p-%p\n", t, &t, t)
	   	} else {
	   		fmt.Printf("dynamic value of 't' without value  %T %v %v\n", t, t, ok)
	   	}

	   	if t, ok := any.(*Circle); ok {
	   		fmt.Printf("dynamic value of 't' with value  %T %v %v\n", t, t, ok)
	   		fmt.Printf("'t' is %T-%p-%p\n", t, &t, t)
	   	} else {
	   		fmt.Printf("dynamic value of 't' without value  %T %v %v\n", t, t, ok)
	   	} */
}

func main() {
	typeTest()
}
