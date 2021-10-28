package main

import "fmt"

func main() {
	resource := make(map[string][]int)

	resource["a"] = append(resource["a"], 1)
	resource["a"] = append(resource["a"], 2)
	fmt.Printf("%+v\n", resource)
}
