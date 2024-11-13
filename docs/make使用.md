在根目录使用
```bash
make IDL=xx init/update
```

在 /rpc/xx-service 使用
```bash
make init/update
```

增加服务后应该在 根目录 Makefile SERVICES 和 SERVICE_IDL_MAP 添加服务名和IDL名

以后可能会 SERVICE 和 SERVICE_IDL_MAP 合并？（摸鱼了）