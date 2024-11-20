package main

import (
	order "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/order/orderservice"
	"log"
)

func main() {
	svr := order.NewServer(new(OrderServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
