package main

import (
	"context"
	order "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/order"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// SubmitOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) SubmitOrder(ctx context.Context, request *order.SubmitOrderRequest) (resp *order.SubmitOrderResponse, err error) {
	// TODO: Your code here...
	return
}

// PayOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PayOrder(ctx context.Context, request *order.PayOrderRequest) (resp *order.PayOrderResponse, err error) {
	// TODO: Your code here...
	return
}

// CancelOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CancelOrder(ctx context.Context, request *order.CancelOrderRequest) (resp *order.CancelOrderResponse, err error) {
	// TODO: Your code here...
	return
}

// GetOrders implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrders(ctx context.Context, request *order.GetOrdersRequest) (resp *order.GetOrdersResponse, err error) {
	// TODO: Your code here...
	return
}

// GetOrderById implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrderById(ctx context.Context, request *order.GetOrderByIdRequest) (resp *order.GetOrderByIdResponse, err error) {
	// TODO: Your code here...
	return
}
