package rpc

import (
	"log"

	etcd "github.com/kitex-contrib/registry-etcd"
)

func InitAll() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	InitCart(r)
	InitOrder(r)
	InitProduct(r)
	InitUser(r)
}
