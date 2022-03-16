# 108

validator

gin 框架使用的 validator
对于 required tag

> This validates that the value is not the data types default zero value. For numbers ensures value is not zero. For strings ensures value is not "". For slices, maps, pointers, interfaces, channels and functions ensures the value is not nil.



不能区分接口传过来的是字段是零值还是没有这个字段,例如

```json
{
	"number":0
}
```

和

```json
{
    
}
```

在转换为结构体`Validator`时，是一样的。

```go
type Validator struct {
	Number    int    `json:"number"`
}
```



如果定义`number`字段为必填，且其值可以为0，如果`number`字段缺失，是检测不出来的。

解决方案有如下两种：

1. 使用指针的形式定义数据，例如

```go
type Validator struct {
	Number    int    `json:"number" binding:"required"`
	NumberPtr *int   `json:"numberPtr" binding:"required"`
}
```

如果`numberPtr`字段缺失，`NumberPtr`则为`nil`，这样可以区分空值问题，但传值时增加了复杂度

```go
var valid Validator
var a int
a = 12
valid.NumberPtr = &a
```

2. 另一种方式是在定义接口时，如果该字段可以为零值，就不定义为必填字段，设置时先读取原值。然后再将json数据转换到该值，缺失的字段保持原值不变。






## reference

1. [gin 框架参数零值 json 绑定的问题](https://blog.csdn.net/weixin_42279809/article/details/107800081)
