package main

import (
	user "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/user/userservice"
	"log"
)

func main() {
	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
