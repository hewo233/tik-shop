update:
	hz update -idl ../idl/api/$(service).thrift
user:
	make service=user update
cart:
	make service=cart update
product:
	make service=product update
order:
	make service=order update
auth:
	make service=auth update

all:
	make user
	make cart
	make product
	make order
	make auth
init:
	hz new -module github.com/hewo/tik-shop/route -idl ../idl/api/auth.thrift
	make user
	make cart
	make product
	make order
# then shall manually delete go.mod go.sum and .gitignore and run go mod tidy at project root after init

build:
	sh ./build.sh

run:
	make build
	sh ./output/bootstrap.sh

.PHONY: update user cart product order auth all init
