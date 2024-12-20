// Code generated by hertz generator. DO NOT EDIT.

package product

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	product "github.com/hewo/tik-shop/route/biz/handler/hewo/tikshop/route/product"
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
		_api.GET("/products", append(_getproductsMw(), product.GetProducts)...)
		_api.POST("/products", append(_createproductMw(), product.CreateProduct)...)
		{
			_product := _api.Group("/product", _productMw()...)
			_product.DELETE("/:id", append(_deleteproductMw(), product.DeleteProduct)...)
			_product.GET("/:id", append(_getproductMw(), product.GetProduct)...)
			_product.PUT("/:id", append(_updateproductMw(), product.UpdateProduct)...)
		}
	}
}
