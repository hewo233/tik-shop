# 根目录 Makefile

# 定义所有的服务目录和对应的 IDL 文件
SERVICES := user-service product-service order-service cart-service

# 定义服务和 IDL 文件的映射
SERVICE_IDL_MAP := user-service=user product-service=product order-service=order cart-service=cart
.PHONY: init update clean

# 初始化生成代码
init:
ifndef IDL
	$(error "请使用 'make IDL=服务名|all init' 指定要生成的服务")
endif
	@if [ "$(IDL)" = "all" ]; then \
		for service in $(SERVICES); do \
			idl_name=$$(echo $(SERVICE_IDL_MAP) | tr ' ' '\n' | grep $$service= | cut -d'=' -f2); \
			echo "Generating code for $$service with IDL $$idl_name.thrift"; \
			$(MAKE) -C ./rpc/$$service init; \
		done; \
	else \
		found=0; \
		for entry in $(SERVICE_IDL_MAP); do \
			service=$${entry%=*}; \
			idl_name=$${entry#*=}; \
			if [ "$(IDL)" = "$$idl_name" ]; then \
				echo "Generating code for $$service with IDL $$idl_name.thrift"; \
				$(MAKE) -C ./rpc/$$service init; \
				found=1; \
				break; \
			fi; \
		done; \
		if [ $$found -eq 0 ]; then \
			echo "未找到与 IDL=$(IDL) 对应的服务"; \
			exit 1; \
		fi; \
	fi

# 更新生成代码
update:
ifndef IDL
	$(error "请使用 'make IDL=服务名|all update' 指定要更新的服务")
endif
	@if [ "$(IDL)" = "all" ]; then \
		for service in $(SERVICES); do \
			idl_name=$$(echo $(SERVICE_IDL_MAP) | tr ' ' '\n' | grep $$service= | cut -d'=' -f2); \
			echo "Updating code for $$service with IDL $$idl_name.thrift"; \
			$(MAKE) -C ./rpc/$$service update; \
		done; \
	else \
		found=0; \
		for entry in $(SERVICE_IDL_MAP); do \
			service=$${entry%=*}; \
			idl_name=$${entry#*=}; \
			if [ "$(IDL)" = "$$idl_name" ]; then \
				echo "Updating code for $$service with IDL $$idl_name.thrift"; \
				$(MAKE) -C ./rpc/$$service update; \
				found=1; \
				break; \
			fi; \
		done; \
		if [ $$found -eq 0 ]; then \
			echo "未找到与 IDL=$(IDL) 对应的服务"; \
			exit 1; \
		fi; \
	fi

# 清理生成的代码
clean:
	@for service in $(SERVICES); do \
		$(MAKE) -C $$service clean; \
	done
