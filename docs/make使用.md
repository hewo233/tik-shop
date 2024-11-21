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

使用
```bash
make clean
make clean-all
```

**不建议使用 clean!!!!**

clean 删除 kitex-gen
clean-all 删除 除了 handler 和 main 之外的所有文件（防止彻底炸掉）