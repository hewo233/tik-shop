package rpc

import (
	"log"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/user/userservice"
)

var UserClient userservice.Client

func InitUser(r discovery.Resolver) {
	c, err := userservice.NewClient("github.com.hewo.tik-shop.user", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	UserClient = c
}
