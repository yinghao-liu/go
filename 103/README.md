# 103

gorm test

## GORM 名称约定

> GORM 倾向于约定，而不是配置。默认情况下，GORM 使用 ID 作为主键，使用结构体名的 `蛇形复数` 作为表名，字段名的 `蛇形` 作为列名，并使用 CreatedAt、UpdatedAt 字段追踪创建、更新时间遵循 GORM 已有的约定，可以减少您的配置和代码量。如果约定不符合您的需求，GORM 允许您自定义配置它们

```go
// 用户模型
type User struct {
	ID     uint   // gorm default id
	Name   string // gorm default name
	Age    int    // gorm default age
	Gender string // gorm default gender

	UserDAO string // gorm default user_dao
	DAOUser string // gorm default dao_user

	UserAAA string //gorm default user_aaa
	AAAUser string //gorm default aaa_user

	UserAAIPA string `gorm:"column:user_aaipa"` //gorm default user_aa_ip_a, change to user_aaipa through column
	AAIPAUser string //gorm default aa_ip_a_user
}
```

对于字段的蛇形转换，如果模型字段中有大写的专有名词，如`IP`,会被单独转换，如果不符合期望可以使用`column`指定列名

## 关联

其中 Belongs To 和 Has One 都是一对一的关系，都是以 user 为主体，不同之处在于它们和关联对象的关系，`Belongs To`的 user 和关联对象（compony）的关系是`属于`，`Has One`的 user 和关联对象（compony）的关系是`属于`,都是站在 user 的角度看的。

`Belongs To`和 SQL 里的一对一是一个概念

### Belongs To（一对一）

### Has One (一对一)

### Has Many （一对多）

### Many To Many （多对多）



## 兼容性

例如：一个字段原来是string类型，之后修改为int类型的处理



## 使用

go test -v -run xxx

## reference

1. [GORM 约定](https://gorm.io/zh_CN/docs/conventions.html)

2. [关联-has_many](https://gorm.io/zh_CN/docs/has_many.html)
