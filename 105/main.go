package main

import (
	"fmt"
	"reflect"
)

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

	fmt.Printf("%#v, %v\n", value, typ)
	fmt.Printf("---------------\n")
	fmt.Printf("value+v:%+v\n, value#v:%#v\n, type:%T\n", secret, secret, secret)
	// alternative:
	//typ := value.Type()  // main.NotknownType
	knd := value.Kind() // struct
	fmt.Printf("%#v, %v\n", knd, knd.String())

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
	testReflect()
}
