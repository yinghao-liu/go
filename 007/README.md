# 007

function

命名返回值重复赋值问题

## 变长参数

variable arguments

变长参数的函数很少会写，但是却经常被使用，例如Printf类的都是使用变长参数表示的。

go的变长参数使用的是`...type`的形式，其对应的参数其实是一个切片。

```go
func Vararg(a string, args ...int) {
	fmt.Printf("a is %T,args is %T\n", a, args) // a is string,args is []int
}
```

那么，它和切片作为参数的函数有什么区别呢？例如：

```go
func SliceArgs(a string, args []int) {
	fmt.Printf("a is %T,args is %T\n", a, args) // a is string,args is []int
}
```

区别是在调用方，传入的参数类型不同。

```go
func TestVarArg(t *testing.T) {
	var slc = []int{1, 2, 3}
	VarArgs("aa", slc)   // cannot use slc (variable of type []int) as int value in argument to VarArgs
	SliceArgs("aa", slc) // OK

	VarArgs("aa", 1, 2, 3)   // ok
	SliceArgs("aa", 1, 2, 3) // too many arguments in call to SliceArgs	have(string, number, number, number) want(string, []int)

	VarArgs("aa", slc...) // OK
}
```

变长参数只接受**单个**、**变长**的指定的类型参数，最后的`slc...`，表述如下

> If the final argument is assignable to a slice type `[]T` and is followed by `...`, it is passed unchanged as the value for a `...T` parameter. In this case no new slice is created.



## reference

1. [函数参数与返回值](https://gitee.com/yinghao-liu/the-way-to-go_ZH_CN/blob/master/eBook/06.2.md)
1. [Passing_arguments_to_..._parameters](https://golang.google.cn/ref/spec#Passing_arguments_to_..._parameters)
