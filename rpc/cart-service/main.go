package main

import (
	cart "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/cart/cartservice"
	"log"
)

func main() {
	svr := cart.NewServer(new(CartServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
