package main

import (
	"context"
	"fmt"

	product "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/product"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct{}

// GetProducts implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProducts(ctx context.Context, request *product.GetProductsRequest) (resp *product.GetProductsReqsponse, err error) {
	// TODO: Your code here...
	resp = &product.GetProductsReqsponse{
		Products: []*product.Product{
			{
				Id:          1,
				Name:        "Product 1",
				Price:       99.99,
				Stock:       100,
				Description: "This is product 1",
			},
			{
				Id:          2,
				Name:        "Product 2",
				Price:       199.99,
				Stock:       200,
				Description: "This is product 2",
			},
		},
	}
	return resp, nil
}

// GetProductById implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProductById(ctx context.Context, request *product.GetProductByIdRequest) (resp *product.GetProductByIdResponse, err error) {
	// TODO: Your code here...
	resp = &product.GetProductByIdResponse{
		Product: &product.Product{
			Id:          request.Id,
			Name:        "Product " + fmt.Sprint(request.Id),
			Price:       123.45,
			Stock:       50,
			Description: "This is the product with ID " + fmt.Sprint(request.Id),
		},
	}
	return resp, nil
}

// CreateProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) CreateProduct(ctx context.Context, request *product.CreateProductRequest) (resp *product.CreateProductResponse, err error) {
	// TODO: Your code here...
	resp = &product.CreateProductResponse{
		Message:   "Product created successfully",
		ProductId: 123, // Fake ID for the new product
	}
	return
}

// UpdateProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, request *product.UpdateProductRequest) (resp *product.UpdateProductResponse, err error) {
	// TODO: Your code here...
	resp = &product.UpdateProductResponse{
		Message: "Product updated successfully",
	}
	return
}

// DeleteProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteProduct(ctx context.Context, request *product.DeleteProductRequset) (resp *product.DeleteProductResponse, err error) {
	// TODO: Your code here...
	resp = &product.DeleteProductResponse{
		Message: "Product deleted successfully",
	}
	return
}
