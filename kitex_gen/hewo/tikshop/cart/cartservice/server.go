// Code generated by Kitex v0.11.3. DO NOT EDIT.
package cartservice

import (
	server "github.com/cloudwego/kitex/server"
	cart "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/cart"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler cart.CartService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler cart.CartService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
