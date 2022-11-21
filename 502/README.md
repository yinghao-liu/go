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

`T.Log`函数



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

benchmark的输入如下

```sh
$ go test -run ^$ -bench BenchmarkRandInt -benchmem
goos: windows
goarch: amd64
pkg: gotest
cpu: Intel(R) Core(TM) i7-8500 CPU @ 3.80GHz
BenchmarkRandInt-6      19528102                53.73 ns/op            8 B/op          1 allocs/op
PASS
ok      gotest  1.317s
```

其中的最终统计信息

```sh
BenchmarkRandInt-6      19528102                53.73 ns/op            8 B/op          1 allocs/op
```

分别是Benchmark名、总执行次数、每次执行耗时、每次执行使用堆内存、每次执行申请内存次数







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





## go test

### 用法

```shell
go test [build/test flags] [packages] [build/test flags & test binary flags]
```



### 选项

常见的选项如下

| 选项          | 含义                                                         |
| ------------- | ------------------------------------------------------------ |
| -v            | 显示更多详细的信息                                           |
| -bench regexp | 默认情况下，benchmark是不运行的<br />该选项指定了一个正则表达式，相符的benchmark会运行，如果要运行所有的<br />benchmark，使用'-bench .'或'-bench=.' |
| -run regexp   | 只运行符合该正则表达式的                                     |
|               |                                                              |
| -cover        | 启用覆盖率分析<br />go test之后会出现覆盖率统计，例如<br />`coverage: 100.0% of statements` |
|               |                                                              |

以下选项可以在test过程中剖析相关信息

| 选项                    | 含义                           |
| ----------------------- | ------------------------------ |
| -benchmem               | benchmark过程中的内存分配统计  |
| -coverprofile cover.out | 将测试覆盖数据写到指定的文件里 |
| -cpuprofile cpu.out     | 将CPU数据写到指定的文件里      |
| -memprofile mem.out     | 将内存分配数据写到指定的文件里 |
|                         |                                |
|                         |                                |
|                         |                                |
|                         |                                |
|                         |                                |





## reference

1. [testing](https://pkg.go.dev/testing)
2. [cmd go test](https://golang.google.cn/cmd/go/)
