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
	items, err := superquery.GetCart(request)
	if err != nil {
		log.Println("Error getting cart items:", err)
		return nil, fmt.Errorf("failed to retrieve cart items: %w", err)
	}
	resp = &cart.GetCartResponse{
		Items: items,
	}
	return
}

// AddToCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddToCart(ctx context.Context, request *cart.AddToCartRequest) (resp *cart.AddToCartResponse, err error) {
	resp, err = superquery.AddToCart(request)
	if err != nil {
		log.Println("Error adding cart items:", err)
		return nil, fmt.Errorf("failed to add to cart: %w", err)
	}
	return
}

// UpdateCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) UpdateCart(ctx context.Context, request *cart.UpdateCartRequest) (resp *cart.UpdateCartResponse, err error) {
	resp, err = superquery.UpdateCart(request)
	if err != nil {
		log.Println("Error updating cart:", err)
		return nil, fmt.Errorf("failed to update cart: %w", err)
	}
	return
}

// RemoveFromCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) RemoveFromCart(ctx context.Context, request *cart.RemoveFromCartRequest) (resp *cart.RemoveFromCartResponse, err error) {

	resp, err = superquery.RemoveFromCart(request)
	if err != nil {
		log.Println("Error removing cart items:", err)
		return nil, fmt.Errorf("failed to remove from cart: %w", err)
	}
	return
}

// ClearCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) ClearCart(ctx context.Context, request *cart.ClearCartRequest) (resp *cart.ClearCartResponse, err error) {

	resp, err = superquery.ClearCart(request)
	if err != nil {
		log.Println("Error clearing cart:", err)
		return nil, fmt.Errorf("failed to clear cart: %w", err)
	}
	return
}

// UpdateQuantity implements the CartServiceImpl interface.
func (s *CartServiceImpl) UpdateQuantity(ctx context.Context, req *cart.UpdateQuantityRequest) (resp *cart.UpdateQuantityResponse, err error) {
	// TODO: Your code here...
	return
}

// ToggleSelect implements the CartServiceImpl interface.
func (s *CartServiceImpl) ToggleSelect(ctx context.Context, req *cart.ToggleSelectRequest) (resp *cart.ToggleSelectResponse, err error) {
	// TODO: Your code here...
	return
}

// RemoveItems implements the CartServiceImpl interface.
func (s *CartServiceImpl) RemoveItems(ctx context.Context, req *cart.RemoveItemsRequest) (resp *cart.RemoveItemsResponse, err error) {
	// TODO: Your code here...
	return
}
