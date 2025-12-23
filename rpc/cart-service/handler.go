package main

import (
	"context"
	"fmt"
	"github.com/hewo/tik-shop/db/model"
	cart "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/cart"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/product"
	"github.com/jinzhu/copier"
	"log"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct {
	CartSQLManage
}

type CartSQLManage interface {
	GetCart(customerID int64) ([]*model.CartItem, error)
	AddToCart(customerID int64, productID int64, quantity int64) (int64, error)
	UpdateQuantity(cartItemID int64, customerID int64, newQuantity int64) error
	ToggleSelect(cartItemIDs []int64, customerID int64, selected bool) error
	RemoveItems(cartItemIDs []int64, customerID int64) error
	ClearCart(customerID int64) error
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, request *cart.GetCartRequest) (resp *cart.GetCartResponse, err error) {
	items, err := s.CartSQLManage.GetCart(request.CustomerId)
	if err != nil {
		log.Println("Error getting cart items:", err)
		return nil, fmt.Errorf("failed to get cart: %w", err)
	}
	groupsMap := make(map[int64]*cart.MerchantGroup)

	for _, item := range items {
		merchantID := item.MerchantID
		cartItem := &cart.CartItem{}
		if err = copier.Copy(cartItem, item); err != nil {
			log.Println("Error copying cart item:", err)
			return nil, fmt.Errorf("failed to process cart item: %w", err)
		}
		cartItem.Product = &product.Product{}
		if err = copier.Copy(cartItem.Product, item.Product); err != nil {
			log.Println("Error copying product:", err)
			return nil, fmt.Errorf("failed to process product: %w", err)
		}
		if _, exists := groupsMap[merchantID]; !exists {
			groupsMap[merchantID] = &cart.MerchantGroup{}
		}
		groupsMap[merchantID].Items = append(groupsMap[merchantID].Items, cartItem)
		if cartItem.Selected {
			groupsMap[merchantID].Subtotal += cartItem.Quantity * cartItem.Product.Price
		}

	}

	resp = &cart.GetCartResponse{}
	for merchantID, group := range groupsMap {
		group.MerchantId = merchantID
		resp.Groups = append(resp.Groups, group)
		resp.TotalSelected += group.Subtotal
	}
	return
}

// AddToCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddToCart(ctx context.Context, request *cart.AddToCartRequest) (resp *cart.AddToCartResponse, err error) {
	cartID, err := s.CartSQLManage.AddToCart(request.CustomerId, request.ProductId, request.Quantity)
	if err != nil {
		log.Println("Error adding to cart:", err)
		return nil, fmt.Errorf("failed to add to cart: %w", err)
	}
	resp = &cart.AddToCartResponse{
		CartItemId: cartID,
	}
	return resp, nil
}

// UpdateQuantity implements the CartServiceImpl interface.
func (s *CartServiceImpl) UpdateQuantity(ctx context.Context, req *cart.UpdateQuantityRequest) (resp *cart.UpdateQuantityResponse, err error) {
	err = s.CartSQLManage.UpdateQuantity(req.CartItemId, req.CustomerId, req.Quantity)
	if err != nil {
		log.Println("Error updating cart item quantity:", err)
		return nil, fmt.Errorf("failed to update quantity: %w", err)
	}
	resp = &cart.UpdateQuantityResponse{
		Success: true,
	}
	return resp, nil
}

// ToggleSelect implements the CartServiceImpl interface.
func (s *CartServiceImpl) ToggleSelect(ctx context.Context, req *cart.ToggleSelectRequest) (resp *cart.ToggleSelectResponse, err error) {
	err = s.CartSQLManage.ToggleSelect(req.CartItemIds, req.CustomerId, req.Selected)
	if err != nil {
		log.Println("Error toggling cart item selection:", err)
		return nil, fmt.Errorf("failed to toggle selection: %w", err)
	}
	resp = &cart.ToggleSelectResponse{
		Success: true,
	}
	return resp, nil
}

// RemoveItems implements the CartServiceImpl interface.
func (s *CartServiceImpl) RemoveItems(ctx context.Context, req *cart.RemoveItemsRequest) (resp *cart.RemoveItemsResponse, err error) {
	err = s.CartSQLManage.RemoveItems(req.CartItemIds, req.CustomerId)
	if err != nil {
		log.Println("Error removing cart items:", err)
		return nil, fmt.Errorf("failed to remove items: %w", err)
	}
	resp = &cart.RemoveItemsResponse{
		Success: true,
	}
	return resp, nil
}

// ClearCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) ClearCart(ctx context.Context, req *cart.ClearCartRequest) (resp *cart.ClearCartResponse, err error) {

	err = s.CartSQLManage.ClearCart(req.CustomerId)
	if err != nil {
		log.Println("Error clearing cart:", err)
		return nil, fmt.Errorf("failed to clear cart: %w", err)
	}
	resp = &cart.ClearCartResponse{
		Success: true,
	}
	return resp, nil
}
