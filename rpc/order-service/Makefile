# order-service/Makefile

# 服务名称
SERVICE_NAME := order

SERVICE_DIR := $(shell pwd)
ROOT_DIR := $(SERVICE_DIR)/../..

# IDL 文件路径（相对于服务目录）
IDL_FILE := $(ROOT_DIR)/idl/$(SERVICE_NAME).thrift

# Go module 名称
MODULE := github.com/hewo/tik-shop

# Kitex 生成代码的输出目录
GEN_DIR := github.com/hewo/tik-shop/kitex_gen

# Kitex 生成代码的选项
MODULE_OPTIONS := -module $(MODULE)
KITEX_OPTIONS := $(MODULE_OPTIONS) -service $(SERVICE_NAME) -use $(GEN_DIR) -verbose

.PHONY: init update clean

# 初始化生成代码（初次生成）
init:
	cd $(ROOT_DIR) && kitex $(MODULE_OPTIONS) $(IDL_FILE)
	kitex $(KITEX_OPTIONS) $(IDL_FILE)

# 更新生成代码（IDL 文件更新后）
update:
	cd $(ROOT_DIR) && kitex $(MODULE_OPTIONS) $(IDL_FILE)
	kitex $(KITEX_OPTIONS) $(IDL_FILE)

# 清理生成的代码
clean:
	rm -rf $(GEN_DIR)

clean-force:
	rm -rf $(GEN_DIR)
	rm -rf ./script
	rm -f ./build.sh ./go.mod ./kitex_info.yaml ./go.sum
