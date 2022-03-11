package main

import "fmt"

func fun1(int, string) error {
	fmt.Printf("in fun1\n")
	return nil
}
func fun2(int, string) (bool, error) {
	fmt.Printf("in fun2\n")
	return true, nil
}
func fun3(int, string) {
	fmt.Printf("in fun3\n")

}

func load(f func(int, string)) {
	fmt.Printf("load fun type: %T\n", f)
}
func load1(f func(int, string) error) {
	fmt.Printf("load fun type: %T\n", f)
}
func main() {
	fmt.Printf("fun1 type: %T\n", fun1)
	fmt.Printf("fun2 type: %T\n", fun2)

	//load(fun1) // cannot use fun1 (type func(int, string) error) as type func(int, string) in argument to load
	//load(fun2) // cannot use fun2 (type func(int, string) (bool, error)) as type func(int, string) in argument to load
	load(fun3)

	load1(fun1)

}
