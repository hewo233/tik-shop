package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hewo/tik-shop/db/model"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/base"
	product "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/product"
	"github.com/jinzhu/copier"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct {
	ProductSqlManage
}

type ProductSqlManage interface {
	CreateProduct(product *model.Product) (productID int64, err error)
	GetProductByID(id int64) (productRet *model.Product, err error)
	ListProducts(merchantID int64, offset int, limit int) (products []*model.Product, err error)
	CheckAndGetProduct(productID int64, merchantID int64) (*model.Product, error)
	UpdateProductByID(product *model.Product) error
	DeleteProductByID(productID int64) (err error)
}

// CreateProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) CreateProduct(ctx context.Context, request *product.CreateProductRequest) (resp *product.CreateProductResponse, err error) {
	p := &model.Product{}

	err = copier.Copy(p, request)
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	p.Status = 1 // 默认上架状态

	id, err := s.ProductSqlManage.CreateProduct(p)

	if err != nil {
		return nil, err
	}
	resp = &product.CreateProductResponse{
		ProductId: id,
	}
	return resp, nil
}

// GetProductByID implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProductByID(ctx context.Context, req *product.GetProductByIDRequest) (resp *product.GetProductByIDResponse, err error) {
	pro, err := s.ProductSqlManage.GetProductByID(req.ProductId)
	if err = copier.Copy(resp, pro); err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}

	if err = copier.Copy(req, pro); err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}

	return resp, nil
}

// ListProducts implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsRequest) (resp *product.ListProductsResponse, err error) {
	offset := int((req.Page - 1) * req.PageSize)
	limit := int(req.PageSize)

	products, err := s.ProductSqlManage.ListProducts(req.MerchantId, offset, limit)
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}

	resp = &product.ListProductsResponse{
		Products: make([]*product.Product, len(products)),
	}
	for i, pro := range products {
		p := &product.Product{}
		err = copier.Copy(p, pro)
		if err != nil {
			return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
		}
		resp.Products[i] = p
	}

	return resp, nil
}

// UpdateProductByID implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateProductByID(ctx context.Context, req *product.UpdateProductByIDRequest) (resp *product.UpdateProductByIDResponse, err error) {
	if req.MerchantId <= 0 || req.ProductId <= 0 {
		return nil, &base.ErrorResponse{Code: consts.StatusBadRequest, Message: "invalid merchant id or product id"}
	}

	existed, err := s.ProductSqlManage.CheckAndGetProduct(req.MerchantId, req.ProductId)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		existed.Name = *req.Name
	}
	if req.Description != nil {
		existed.Description = *req.Description
	}
	if req.Price != nil {
		existed.Price = *req.Price
	}
	if req.Status != nil {
		existed.Status = *req.Status
	}

	err = s.ProductSqlManage.UpdateProductByID(existed)
	if err != nil {
		return nil, err
	}

	resp = &product.UpdateProductByIDResponse{}
	err = copier.Copy(resp, existed)
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}

	return resp, nil

}

// DeleteProductByID implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteProductByID(ctx context.Context, req *product.DeleteProductByIDRequest) (resp *product.DeleteProductByIDResponse, err error) {
	if req.MerchantId <= 0 || req.ProductId <= 0 {
		return nil, &base.ErrorResponse{Code: consts.StatusBadRequest, Message: "invalid merchant id or product id"}
	}

	_, err = s.ProductSqlManage.CheckAndGetProduct(req.MerchantId, req.ProductId)
	if err != nil {
		return nil, err
	}

	err = s.ProductSqlManage.DeleteProductByID(req.ProductId)
	if err != nil {
		return nil, err
	}

	resp = &product.DeleteProductByIDResponse{
		Success: true,
	}

	return resp, nil
}

// ModifyStockByID implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ModifyStockByID(ctx context.Context, req *product.ModifyStockByIDRequest) (resp *product.ModifyStockByIDResponse, err error) {
	if req.MerchantId <= 0 || req.ProductId <= 0 {
		return nil, &base.ErrorResponse{Code: consts.StatusBadRequest, Message: "invalid merchant id or product id"}
	}

	existed, err := s.ProductSqlManage.CheckAndGetProduct(req.MerchantId, req.ProductId)
	if err != nil {
		return nil, err
	}

	if existed.Stock != req.CurrentStock {
		return nil, &base.ErrorResponse{Code: consts.StatusConflict, Message: "stock mismatch"}
	}

	existed.Stock += req.Delta
	err = s.ProductSqlManage.UpdateProductByID(existed)
	if err != nil {
		return nil, err
	}

	resp = &product.ModifyStockByIDResponse{
		Stock: existed.Stock,
	}

	return resp, nil
}
