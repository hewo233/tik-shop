package main

import (
	"github.com/hewo/tik-shop/db/connectDB"
	"github.com/hewo/tik-shop/db/query"
	"github.com/hewo/tik-shop/db/superquery"
	"log"
	"net"

	order "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/order/orderservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
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

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8891")
	svr := order.NewServer(&(OrderServiceImpl{
		OrderSqlManage: superquery.NewOrderSqlManageImpl(),
	}),
		server.WithServiceAddr(addr),
		// 指定 Registry 与服务基本信息
		server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "github.com.hewo.tik-shop.order",
			},
		),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
