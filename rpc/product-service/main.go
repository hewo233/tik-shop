package main

import (
	"github.com/hewo/tik-shop/db/connectDB"
	"github.com/hewo/tik-shop/db/query"
	"github.com/hewo/tik-shop/db/superquery"
	"github.com/hewo/tik-shop/shared/consts"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	product "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/product/productservice"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {

	database, err := connectDB.ConnectDB(consts.RpcDBEnvPath)
	if err != nil {
		log.Println(err)
	}
	query.SetDefault(database)

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8892")
	svr := product.NewServer(&ProductServiceImpl{
		ProductSqlManage: superquery.NewProductSqlManageImpl(),
	},
		server.WithServiceAddr(addr),
		// 指定 Registry 与服务基本信息
		server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "github.com.hewo.tik-shop.product",
			},
		),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
