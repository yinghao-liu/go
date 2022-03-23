# 502

gotest

`go test`会重新单独编译符合"\*\_test.go"模式的文件，这些文件会包含`test`函数、`benchmark`函数、`fuzz`测试（1.18）和`example`函数，更多信息见'go help testfunc'，Each listed package causes the execution of a separate test binary，以`_`或`.`开头的文件会被忽略。

以"\_test"结尾的包声明的测试文件会被单独编译，并随主测试程序运行，正常一个文件夹里只能有一个包（package）,但是允许存在以`_test`结尾的其他包名。

go tool 会忽略名为"testdata"的文件夹，这可以帮助搞定 test 需要的辅助数据



## Test

test函数是`go test`默认会执行的，其格式：

```go
func TestXxx(*testing.T)
```

一个test函数即可作为一个测试用例。

在test函数里，使用`T`的`Error`、`Errorf`、`Fail`、`FailNow`、`Failed`、`Fatal`、`Fatalf`等函数标记失败。





## Benchmark

benchmark函数的格式：

```go
func BenchmarkXxx(*testing.B)
```

例如

```go
func BenchmarkRandInt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        rand.Int()
    }
}
```

benchmark函数必须执行`b.N`次，它能够自适应以被察觉到的次数。





## Example

example函数是`go test`默认会执行的，其格式：

```go
func Example()
```

或

```go
func ExampleHello()
```

例如

```go
func ExampleHello() {
    fmt.Println("hello")
    // Output: hello
}
```





## Fuzzing 

fuzz测试的格式：

```go
func FuzzXxx(*testing.F)
```











## reference

1. [testing](https://pkg.go.dev/testing)
2. [cmd go test](https://golang.google.cn/cmd/go/)
