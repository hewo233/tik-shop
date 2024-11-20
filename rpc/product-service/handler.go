package main

import (
	"context"
	product "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/product"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct{}

// GetProducts implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProducts(ctx context.Context, request *product.GetProductsRequest) (resp *product.GetProductsReqsponse, err error) {
	// TODO: Your code here...
	return
}

// GetProductById implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProductById(ctx context.Context, requset *product.GetProductByIdRequest) (resp *product.GetProductByIdResponse, err error) {
	// TODO: Your code here...
	return
}

// CreateProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) CreateProduct(ctx context.Context, request *product.CreateProductRequest) (resp *product.CreateProductResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, request *product.UpdateProductRequest) (resp *product.UpdateProductResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteProduct(ctx context.Context, request *product.DeleteProductRequset) (resp *product.DeleteProductResponse, err error) {
	// TODO: Your code here...
	return
}
