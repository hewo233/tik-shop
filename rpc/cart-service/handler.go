package main

import (
	"context"
	cart "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/cart"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, request *cart.GetCartRequest) (resp *cart.GetCartResponse, err error) {
	// TODO: Your code here...
	return
}

// AddToCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddToCart(ctx context.Context, request *cart.AddToCartRequest) (resp *cart.AddToCartResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) UpdateCart(ctx context.Context, request *cart.UpdateCartRequest) (resp *cart.UpdateCartResponse, err error) {
	// TODO: Your code here...
	return
}

// RemoveFromCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) RemoveFromCart(ctx context.Context, request *cart.RemoveFromCartRequest) (resp *cart.RemoveFromCartResponse, err error) {
	// TODO: Your code here...
	return
}

// ClearCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) ClearCart(ctx context.Context, request *cart.ClearCartRequest) (resp *cart.ClearCartResponse, err error) {
	// TODO: Your code here...
	return
}
