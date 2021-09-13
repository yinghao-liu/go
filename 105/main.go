package main

import (
	"fmt"
	"math"
	"reflect"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

// variable to investigate:
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func testReflect() {
	value := reflect.ValueOf(secret) // <main.NotknownType Value>
	typ := reflect.TypeOf(secret)    // main.NotknownType
	// alternative:
	//typ := value.Type()  // main.NotknownType
	fmt.Println(typ)
	knd := value.Kind() // struct
	fmt.Println(knd)

	// iterate through the fields of the struct:
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		// error: panic: reflect.Value.SetString using value obtained using unexported field
		//value.Field(i).SetString("C#")
	}

	// call the first method, which is String():
	results := value.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]
}

func main() {
	var any interface{}

	fmt.Printf("'any' is %T-%p-%p\n", any, &any, any)
	sq1 := new(Square)
	any = sq1
	fmt.Printf("'any' is %T-%p-%p\n", any, &any, any)

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
