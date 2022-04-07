# conventions

即使是内核的源码文件夹，其命名也有让人有疑问的地方，比如为什么有的用复数有的用单数，有一些约定是必要的。

![v5.17](image/kernel517.png)

go的一些约定是强制性的，会影响语法，所以有必要在这里做一些总结



172







There's a few guidelines to follow.

1. File names that begin with "." or "_" are ignored by the go tool
2. Files with the suffix `_test.go` are only compiled and run by the `go test` tool.
3. Files with os and architecture specific suffixes automatically follow those same constraints, e.g. `name_linux.go` will only build on linux, `name_amd64.go` will only build on amd64. This is the same as having a `//+build amd64` line at the top of the file





## 名称

### 文件夹名

几个特殊的文件夹

#### internal

`internal`用于保护其下的代码，使其对其他包不可见，只能由它所在的包及其子包使用。

如下，`z`包只能由`domain`包及其子包使用，其他包不可见。详见[参考4][]

![internal](image/internal.png)

#### vendor

Go 1.6支持使用外部依赖的本地拷贝，例如你可以把github上使用的包拷贝到本地。

其使用和`internal`类似，而且使用时不用包含前缀。详见[参考4][]



#### go tool忽略的文件

`testdata`文件夹会被go tool忽略

以`.`和`_`开头的文件和文件夹会被go tool忽略







### 文件名



见[参考1][]

### 包名



### 变量名





## FAQ

1. 函数返回结构体时，如果出错要构造空结构体





## reference

1. [filename conventions #36060][参考1]
2. [what-are-conventions-for-filenames-in-go][参考2]
3. [names](https://golang.google.cn/doc/effective_go#names)
4. [cmd go][参考4]





[参考1]:https://github.com/golang/go/issues/36060
[参考2]:https://stackoverflow.com/questions/25161774/what-are-conventions-for-filenames-in-go
[参考4]:https://golang.google.cn/cmd/go/