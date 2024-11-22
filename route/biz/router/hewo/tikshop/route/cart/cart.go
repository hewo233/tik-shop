// Code generated by hertz generator. DO NOT EDIT.

package cart

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	cart "github.com/hewo/tik-shop/route/biz/handler/hewo/tikshop/route/cart"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		_api.DELETE("/cart", append(_deletecartMw(), cart.DeleteCart)...)
		_api.GET("/cart", append(_getcartMw(), cart.GetCart)...)
		_cart := _api.Group("/cart", _cartMw()...)
		_cart.DELETE("/:productId", append(_deletecartitemMw(), cart.DeleteCartItem)...)
		_cart.PUT("/:productId", append(_updatecartitemMw(), cart.UpdateCartItem)...)
		_api.POST("/cart", append(_addtocartMw(), cart.AddToCart)...)
	}
}