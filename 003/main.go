package main

import (
	"fmt"
	"math"
)

type Square struct {
	Side float32
}

func (sq *Square) Area() float32 {
	return sq.Side * sq.Side
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

	fmt.Printf("'any' is [%T]-[%T]-[%p]-[%p]\n", any, &any, &any, any)
	var any2 interface{}
	if any == any2 {
		fmt.Printf("1---   any == any2: equal\n")
	}

	sq1 := new(Square)
	any = sq1
	fmt.Printf("'any' is [%T]-[%T]-[%p]-[%p]\n", any, &any, &any, any)

	var sq2 Square
	any2 = sq2
	fmt.Printf("'any2' is [%T]-[%p]-[%p]\n", any2, &any2, any2)

	if any == any2 {
		fmt.Printf("2---   any == any2: equal\n")
	}

	var any3 interface{}
	var sq3 Square
	sq3.Side = 111 // without this line, any2 and any3 is equal
	any3 = sq3
	fmt.Printf("'any2' is [%T]-[%p]-[%p]\n", any2, &any2, any2)
	fmt.Printf("'any3' is [%T]-[%p]-[%p]\n", any3, &any3, any3)
	if any2 == any3 {
		fmt.Printf("3---   any2 == any3: equal\n")
	}

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
