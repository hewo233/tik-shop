package main

import (
	"github.com/hewo/tik-shop/db/connectDB"
	"github.com/hewo/tik-shop/db/query"
	"github.com/hewo/tik-shop/db/superquery"
	"github.com/hewo/tik-shop/rpc/user-service/config"
	usinit "github.com/hewo/tik-shop/rpc/user-service/init"
	"github.com/hewo/tik-shop/rpc/user-service/pkg/paseto"
	"github.com/hewo/tik-shop/shared/consts"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	user "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/user/userservice"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {

	database, err := connectDB.ConnectDB(consts.RpcDBEnvPath)
	if err != nil {
		log.Println(err)
	}
	query.SetDefault(database)

	usinit.InitConfig()

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8893")

	tg, err := paseto.NewTokenGenerator(
		config.GlobalUserServerConfig.PasetoInfo.PrivateKey,
		[]byte(config.GlobalUserServerConfig.PasetoInfo.Implicit),
	)
	if err != nil {
		log.Fatal(err)
	}

	svr := user.NewServer(&(UserServiceImpl{
		LoginSqlManage: superquery.NewLoginSqlManageImpl(),
		TokenGenerator: tg,
	}),
		server.WithServiceAddr(addr),
		// 指定 Registry 与服务基本信息
		server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "github.com.hewo.tik-shop.user",
			},
		),
	)

	//svr := user.NewServer(new(UserServiceImpl))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
