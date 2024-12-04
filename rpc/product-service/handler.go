package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hewo/tik-shop/db/model"
	"github.com/hewo/tik-shop/db/superquery"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/base"
	product "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/product"
	"github.com/jinzhu/copier"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct{}

// GetProducts implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProducts(ctx context.Context, request *product.GetProductsRequest) (resp *product.GetProductsReqsponse, err error) {
	products, err := superquery.GetProducts(request.Page, request.Limit)
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	resp = &product.GetProductsReqsponse{
		Products: products,
	}
	return resp, nil
}

// GetProductById implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProductById(ctx context.Context, request *product.GetProductByIdRequest) (resp *product.GetProductByIdResponse, err error) {
	p, err := superquery.GetProductById(request.Id)
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	resp = &product.GetProductByIdResponse{
		Product: p,
	}
	return resp, nil
}

// CreateProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) CreateProduct(ctx context.Context, request *product.CreateProductRequest) (resp *product.CreateProductResponse, err error) {
	p := &model.Product{}
	err = copier.Copy(&p, request)
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	err = superquery.CreateProduct(p)
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	resp = &product.CreateProductResponse{
		Message:   "Product created successfully",
		ProductId: -1, // I mean, maybe we dont need to return id?
	}
	return
}

// UpdateProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, request *product.UpdateProductRequest) (resp *product.UpdateProductResponse, err error) {
	p := &model.Product{}
	err = copier.Copy(&p, request)
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	err = superquery.UpdateProduct(p)
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	resp = &product.UpdateProductResponse{
		Message: "Product updated successfully",
	}
	return
}

// DeleteProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteProduct(ctx context.Context, request *product.DeleteProductRequest) (resp *product.DeleteProductResponse, err error) {
	err = superquery.DeleteProduct(request.Id)
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	resp = &product.DeleteProductResponse{
		Message: "Product deleted successfully",
	}
	return
}
