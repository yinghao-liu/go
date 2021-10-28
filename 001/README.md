# 001

slice
切片本身是一个指针

数组声明时需要指定长度，如`var arr = [5]int`, 即使已经给定了确定的数据也需要使用`...`表示，如`var arrLazy = [...]int{5, 6, 7, 8, 22}`
因为省略[]里长度的形式由切片使用

> var arrAge = [5]int{18, 20, 15, 22, 16}
> 注意 [5]int 可以从左边起开始忽略：[10]int {1, 2, 3} :这是一个有 10 个元素的数组，除了前三个元素外其他元素都为 0。

> var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
> 只有索引 3 和 4 被赋予实际的值，其他元素都被设置为空的字符串

声明切片的格式是： var identifier []type（不需要说明长度）。

一个切片在未初始化之前默认为 nil，长度为 0。

切片的初始化格式是：var slice1 []type = arr1[start:end]。

## reference

- [the-way-to-go-slice](https://github.com/yinghao-liu/the-way-to-go_ZH_CN/blob/master/eBook/07.2.md)
- [microsoft-slice](https://docs.microsoft.com/zh-cn/learn/modules/go-data-types/2-slices?ns-enrollment-type=LearningPath&ns-enrollment-id=learn.languages.go-first-steps)
