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
type ProductServiceImpl struct {
	ProductSqlManage
}

type ProductSqlManage interface {
	CreateProduct(product *model.Product) (productID int64, err error)
	GetProductByID(id int64) (productRet *model.Product, err error)
	ListProducts(merchantID int64, offset int, limit int) (products []*model.Product, err error)
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

// ModifyStock implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ModifyStock(ctx context.Context, req *product.ModifyStockRequest) (resp *product.ModifyStockResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateProductByID implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateProductByID(ctx context.Context, req *product.UpdateProductByIDRequest) (resp *product.UpdateProductByIDResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteProductByID implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteProductByID(ctx context.Context, req *product.DeleteProductByIDRequest) (resp *product.DeleteProductByIDResponse, err error) {
	// TODO: Your code here...
	return
}
