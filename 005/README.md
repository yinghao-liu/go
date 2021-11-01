# 005

1. 类型和作用在它上面定义的方法必须在同一个包里定义, 它们可以存在在不同的源文件里，这就是为什么不能在 int、float 或类似这些的类型上定义方法
2. 别名类型没有原始类型上已经定义过的方法。
3. 当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌，这在效果上等同于外层类型 继承 了这些方法
4. 结构体里一个好的策略是创建一些小的、可复用的类型作为一个工具箱，用于组成域类型。

## 多重继承的重名问题

BaseInfo 包含了 Base1 和 Base2，Base1 和 Base2 都实现了`show()`方法，那么 BaseInfo 不能直接使用 BaseInfo `show()`方法，
因为会产生一个错误`ambiguous selector a.show`，需要明确指定具体要使用哪个

如果 BaseInfo 和 Base1 都定义了`show()`方法，则会使用 BaseInfo 的方法

## reference

1. [方法](https://github.com/yinghao-liu/the-way-to-go_ZH_CN/blob/master/eBook/10.6.md)
