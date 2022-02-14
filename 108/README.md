# 108

validator

gin 框架使用的 validator
对于 required tag

> This validates that the value is not the data types default zero value. For numbers ensures value is not zero. For strings ensures value is not "". For slices, maps, pointers, interfaces, channels and functions ensures the value is not nil.

不能区分接口传过来的是字段是零值还是没有这个字段，虽然可以用指针的形式定义结构体,例如

```go
type Validator struct {
	Number    int    `json:"number" binding:"required"`
	NumberPtr *int   `json:"numberPtr" binding:"required"`
}
```

这样可以区分零值问题，但传值时增加了复杂度

```go
var valid Validator
var a int
a = 12
valid.NumberPtr = &a
```

最终在定义接口时选择了如果该字段可以为零值，就不定义为必填字段。

## reference

1. [gin 框架参数零值 json 绑定的问题](https://blog.csdn.net/weixin_42279809/article/details/107800081)
