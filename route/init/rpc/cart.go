package rpc

import (
	"log"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/cart/cartservice"
)

var CartClient cartservice.Client

func InitCart(r discovery.Resolver) {
	c, err := cartservice.NewClient("github.com.hewo.tik-shop.cart", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	CartClient = c
}
