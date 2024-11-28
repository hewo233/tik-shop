package main

import (
	"github.com/hewo/tik-shop/db/connectDB"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/hewo/tik-shop/db/query"
	cart "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/cart/cartservice"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	database, err := connectDB.ConnectDB()
	if err != nil {
	}
	query.SetDefault(database)
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")
	svr := cart.NewServer(new(CartServiceImpl),
		server.WithServiceAddr(addr),
		// 指定 Registry 与服务基本信息
		server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "github.com.hewo.tik-shop.cart",
			},
		),
	)
	// svr := cart.NewServer(new(CartServiceImpl))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
