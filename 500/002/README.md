# 002

依赖倒置下的分层架构

![depend](image/depend.png)



基础设施层应该只做领域层定义接口的实现，比如具体的持久化。而一些通用的工具函数，可以放在领域层的通用域中。

一个符合依赖倒置的资源库的结构



![repository](image/repository.png)