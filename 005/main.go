package main

import (
	"fmt"
	"reflect"
	st "structtest/struct"
)

// alias
type BaseAlias st.Base2

//
type Name struct {
	Name string
}

func main() {
	var a st.BaseInfo
	a.Name = "abc"
	a.Description = "cc"
	a.Gender = "male"
	//a.unexported = "unexported" //a.unexported undefined (type st.BaseInfo has no field or method unexported)
	a.Show() // if BaseInfo did not declare show(), there is an error: ambiguous selector a.show
	//fmt.Printf("%+v\n", a.base) // a.base undefined (cannot refer to unexported field or method base)

	// alias
	//var b BaseAlias
	//b.show() // b.show undefined (type Base2 has no field or method show)

	var Name Name
	Name.Name = "Name"
	fmt.Printf("Name is %+v\n", Name)
	var intvalue int
	//var c interface{} = &intvalue
	tt := reflect.TypeOf(intvalue)
	fmt.Println(tt)

	/*************************************/
	type S struct {
		F string `species:"gopher" color:"blue"`
	}

	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))

	/***********************/
	type S struct {
		F string `species:"gopher" color:"blue"`
	}

}
