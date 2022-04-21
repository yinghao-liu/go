package main

import "fmt"

func nilMap() {
	var m map[string]string
	fmt.Printf("m is %T\n", m)
	fmt.Printf("m is %v\n", m) // Println == %v
	fmt.Println(m)
	fmt.Printf("m is %p\n", m)
	fmt.Printf("len(m) is %d\n", len(m))
	//m["a"] = "a" //panic: assignment to entry in nil map

	var a map[string]int
	for i, j := range a {
		fmt.Printf("nilmap i:%v,j:%v\n", i, j)
	}
	fmt.Printf("nilmap retrieve\n")
	avalue := a["a"] // ok avalue = 0(int())
	fmt.Printf("nilmap retrieve avalue is %v\n", avalue)
	fmt.Printf("nilmap assignment\n")
	a["a"] = 1 // panic: assignment to entry in nil map
	fmt.Printf("nilmap assignment a is %v\n", a)
}
func MapTest() {
	resource := make(map[string][]int)

	resource["a"] = append(resource["a"], 1)
	resource["a"] = append(resource["a"], 2)
	fmt.Printf("%+v\n", resource)
}
func main() {
	nilMap()
}
