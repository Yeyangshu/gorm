# gorm api学习

## 安装
```go
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## 测试代码
- 上下文测试：context/context_test.go

## gorm方法
链式方法链接：https://github.com/go-gorm/gorm/blob/master/chainable_api.go

终结方法链接：https://github.com/go-gorm/gorm/blob/master/finisher_api.go

# gorm源码
## gorm是协程安全的嘛？
知乎问答：https://www.zhihu.com/question/430806549