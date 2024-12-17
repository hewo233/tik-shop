在 route 下面配一个 config ，补全 init 导入配置

写 common 包， 生成中间件

写中间件，文档

目前做了：

写了 auth 的 common 包，用于生成中间件

修改了 api 的 idl，现在 login 等操作也放在 user 层面

增加了 config 的配置，用来做全局配置(从 config.yaml 去读入)

增加了 shared 目录，放置 errno 和 consts

增加了 key ,用于 paseto

由于我懒，所以本应该不使用 c.String, 而是使用 c.JSON + errno 做标准化错误处理， TODO