# user-service-service/Makefile

# 服务名称
SERVICE_NAME := product

# IDL 文件路径（相对于服务目录）
IDL_FILE := ../../idl/$(SERVICE_NAME).thrift

# Go module 名称
MODULE := github.com/hewo/tik-shop/rpc/$(SERVICE_NAME)-service

# Kitex 生成代码的输出目录
GEN_DIR := ./kitex_gen

# Kitex 生成代码的选项
KITEX_OPTIONS := -module $(MODULE) -service $(SERVICE_NAME) -verbose

.PHONY: init update clean

# 初始化生成代码（初次生成）
init:
	kitex $(KITEX_OPTIONS) $(IDL_FILE)

# 更新生成代码（IDL 文件更新后）
update:
	kitex $(KITEX_OPTIONS) $(IDL_FILE)

# 清理生成的代码
clean:
	rm -rf $(GEN_DIR)

clean-force:
	rm -rf $(GEN_DIR)
	rm -rf ./script
	rm -f ./build.sh ./go.mod ./kitex_info.yaml ./go.sum
