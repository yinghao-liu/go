package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func struct2mapTest(sts map[string]interface{}) {
	var st Student
	st.Name = "francis"
	st.Age = 18
	Struct2Map(st, sts)
}
func Struct2Map(st interface{}, dest map[string]interface{}) error {
	data, e := json.Marshal(st)
	if nil != e {
		return e
	}
	if e := json.Unmarshal(data, &dest); nil != e {
		return e
	}
	return nil
}
func main() {

	sts := make(map[string]interface{})
	sts["gender"] = "male"
	fmt.Printf("before convert:%+v\n", sts)
	struct2mapTest(sts)
	fmt.Printf("after convert:%+v\n", sts)

}
