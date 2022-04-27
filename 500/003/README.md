# 003

六边形架构





![six](image/six.png)







分层参考洋葱架构

![onion](image/onion.png)



`Tests`可以用于在不依赖外部服务的情况下，对业务领域逻辑进行测试，只需要在依赖注入时将真实的服务替换成Tests。

对于像go这种支持包测试的，测试文件和当前包放在一起，但是包名不一样，而且在层级上属于最外层。如下

![hierarchy](image/hierarchy.drawio.png)





## FAQ





## reference

1. [map types](https://golang.google.cn/ref/spec#Map_types)
