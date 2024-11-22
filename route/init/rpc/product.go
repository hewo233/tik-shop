package rpc

import (
	"log"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/product/productservice"
)

var ProductClient productservice.Client

func InitProduct(r discovery.Resolver) {
	c, err := productservice.NewClient("github.com.hewo.tik-shop.product", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	ProductClient = c
}
