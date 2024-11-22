package rpc

import (
	"log"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/order/orderservice"
)

var OrderClient orderservice.Client

func InitOrder(r discovery.Resolver) {
	c, err := orderservice.NewClient("github.com.hewo.tik-shop.order", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	OrderClient = c
	log.Println("Inited OrderClient")
}
