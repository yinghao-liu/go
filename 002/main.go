package main

import "fmt"

/*
For an expression x of interface type and a type T, the primary expression
x.(T)
asserts that x is not nil and that the value stored in x is of type T. The notation x.(T) is called a type assertion.

1. More precisely, if T is not an interface type, x.(T) asserts that the dynamic type of x is identical to the type T.
In this case, T must implement the (interface) type of x; otherwise the type assertion is invalid since it is not possible for x to store a value of type T.

2. If T is an interface type, x.(T) asserts that the dynamic type of x implements the interface T.

3. If the type assertion holds, the value of the expression is the value stored in x and its type is T.
If the type assertion is false, a run-time panic occurs. In other words, even though the dynamic type of x is known only at run time,
the type of x.(T) is known to be T in a correct program.
*/
type A struct {
	a string
	b int
}

func (aa A) m() {
	fmt.Printf("%s\n", "ss")
}

func typeAssertion() {
	/*--------------- case 1 as described above ---------------*/
	var any interface {
		m()
	}

	var aInstance A
	aInstance.a = "test"
	//t1 := any.(int)                // int does not implement interface { "".m() } (missing m method)
	//fmt.Printf("%T, %v\n", t1, t1)

	any = aInstance
	t := any.(A)
	fmt.Printf("case 1 %T, %v\n", t, t) //main.A, {test 0}

	/*----------------case 2 as described above--------------------*/
	type withm interface {
		m()
	}
	//withm = aInstance
	ti := any.(withm)
	fmt.Printf("case 2 %T, %v\n", ti, ti) //main.A, {test 0}

	/*----------------case 3 as described above--------------------*/
	var x interface{} = 7 // x has dynamic type int and value 7
	//i := x.(string)       // panic: interface conversion: interface {} is int, not string
	if i, ok := x.(string); ok {
		fmt.Printf("%T, %v\n", i, i)
	}
	if i, ok := x.(int); ok {
		fmt.Printf("%T, %v\n", i, i)
	}

}

func main() {
	typeAssertion()
}
