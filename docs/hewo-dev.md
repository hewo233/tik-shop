由于我懒，所以本应该不使用 c.String, 而是使用 c.JSON + errno 做标准化错误处理， TODO

后续会修改 Register, 先过一层生成 验证码的，然后在 Register 验证,做一个身份管理。

增加 etcd 的 Retry 等等功能。

目前 UpdateUser 不做验证， TODO