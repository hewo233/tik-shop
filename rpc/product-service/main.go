package main

import (
	product "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/product/productservice"
	"log"
)

func main() {
	svr := product.NewServer(new(ProductServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
