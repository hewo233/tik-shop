# DB约定
## 概览
`init.go`用于初始化数据库，我们先搬用kitex生成的各个微服务的结构体，并增加`gorm`注释，然后再`init.go`分别构造数据库表和生成`query`文件夹

`query`文件夹包含了对数据库的各种标准化操作，通过调用这些标准化操作我们可以更加安全的使用数据库

`superquery`是对`query`的进一步封装，一方面，我们通过对`query`的组装形成更高阶的函数，另一方面，我们把数据库返回的model.XXX转化为各个微服务支持的格式

## 操作指南
首先，我们要分别在各个微服务的`main.go`连接数据库
```	go
        database, err := connectDB.ConnectDB()
	if err != nil {}
	query.SetDefault(database)
```
当你在`superquery`写好函数之后就可以在各个handler里面调用了