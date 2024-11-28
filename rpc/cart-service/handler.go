package main

import (
	"context"
	"fmt"
	"github.com/hewo/tik-shop/db/superquery"
	cart "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/cart"
	"log"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, request *cart.GetCartRequest) (resp *cart.GetCartResponse, err error) {
	// 获取购物车中的所有项
	items, err := superquery.GetCart()
	if err != nil {
		// 记录错误日志
		log.Println("Error getting cart items:", err)

		// 返回带有错误信息的响应
		return nil, fmt.Errorf("failed to retrieve cart items: %w", err)
	}
	resp = &cart.GetCartResponse{
		Items: items,
	}
	return
}

// AddToCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddToCart(ctx context.Context, request *cart.AddToCartRequest) (resp *cart.AddToCartResponse, err error) {
	// TODO: Your code here...
	resp = &cart.AddToCartResponse{
		Message: "Product added to cart successfully",
	}
	return
}

// UpdateCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) UpdateCart(ctx context.Context, request *cart.UpdateCartRequest) (resp *cart.UpdateCartResponse, err error) {
	// TODO: Your code here...
	resp = &cart.UpdateCartResponse{
		Message: "Cart updated successfully",
	}
	return
}

// RemoveFromCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) RemoveFromCart(ctx context.Context, request *cart.RemoveFromCartRequest) (resp *cart.RemoveFromCartResponse, err error) {
	// TODO: Your code here...
	resp = &cart.RemoveFromCartResponse{
		Message: "Product removed from cart successfully",
	}
	return
}

// ClearCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) ClearCart(ctx context.Context, request *cart.ClearCartRequest) (resp *cart.ClearCartResponse, err error) {
	// TODO: Your code here...
	resp = &cart.ClearCartResponse{
		Message: "Cart cleared successfully",
	}
	return
}
