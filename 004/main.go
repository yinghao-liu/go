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

/*
Type switches
A type switch compares types rather than values. It is otherwise similar to an expression switch.
It is marked by a special switch expression that has the form of a type assertion using the keyword type rather than an actual type:

switch x.(type) {
// cases
}
Cases then match actual types T against the dynamic type of the expression x. As with type assertions, x must be of interface type,
and each non-interface type T listed in a case must implement the type of x. The types listed in the cases of a type switch must all be different.
(T can be an interface type)

Instead of a type, a case may use the predeclared identifier nil; that case is selected when the expression in the TypeSwitchGuard is a nil interface value. There may be at most one nil case.

Given an expression x of type interface{}, the following type switch:

switch i := x.(type) {
case nil:
	printString("x is nil")                // type of i is type of x (interface{})
case int:
	printInt(i)                            // type of i is int
case float64:
	printFloat64(i)                        // type of i is float64
case func(int) float64:
	printFunction(i)                       // type of i is func(int) float64
case bool, string:
	printString("type is bool or string")  // type of i is type of x (interface{})
default:
	printString("don't know the type")     // type of i is type of x (interface{})
}
could be rewritten:

v := x  // x is evaluated exactly once
if v == nil {
	i := v                                 // type of i is type of x (interface{})
	printString("x is nil")
} else if i, isInt := v.(int); isInt {
	printInt(i)                            // type of i is int
} else if i, isFloat64 := v.(float64); isFloat64 {
	printFloat64(i)                        // type of i is float64
} else if i, isFunc := v.(func(int) float64); isFunc {
	printFunction(i)                       // type of i is func(int) float64
} else {
	_, isBool := v.(bool)
	_, isString := v.(string)
	if isBool || isString {
		i := v                         // type of i is type of x (interface{})
		printString("type is bool or string")
	} else {
		i := v                         // type of i is type of x (interface{})
		printString("don't know the type")
	}
}
The type switch guard may be preceded by a simple statement, which executes before the guard is evaluated.

The "fallthrough" statement is not permitted in a type switch.
*/
type B struct {
	a string
	b int
}

func (bb B) String() string {
	return fmt.Sprintf("%s\n", "with String")
}

func typeSwitch(value interface{}) {
	type Stringer interface {
		String() string
	}

	type ERRStringer interface {
		Error() string
	}
	//ss := value.(type) // use of .(type) outside type switch     // can only be used in `type switch`
	//fmt.Printf("string %v", ss)

	// type Stringer and B are all match, depend on their order appeared
	switch str := value.(type) {
	case string:
		fmt.Printf("string %v", str)

	case Stringer:
		fmt.Printf("Stringer %v", str.String())
	case B:
		fmt.Printf("B %v", str)
	case ERRStringer:
		fmt.Printf("ERRStringer %v", str.Error())
	default:
		fmt.Printf("default %v, %T", str, str)
	}
}

func main() {
	//typeAssertion()
	//var a error
	var bString B
	typeSwitch(bString)
	var s int
	typeSwitch(s)
}
