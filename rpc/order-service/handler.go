package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/base"
	order "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/order"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct {
	OrderSqlManage
}
type OrderSqlManage interface {
	SubmitOrder(UserId int64, Items []*order.OrderItem, Address *order.Address, PaymentMethod string) (*order.SubmitOrderResponse, error)
	PayOrder(orderId int64, PaymentMethod string, PaymentDetails *order.PaymentDetails) (*order.PayOrderResponse, error)
	CancelOrder(orderId int64) (*order.CancelOrderResponse, error)
	GetOrders(userId int64) (*order.GetOrdersResponse, error)
	GetOrderById(orderId int64) (*order.GetOrderByIdResponse, error)
}

// SubmitOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) SubmitOrder(ctx context.Context, request *order.SubmitOrderRequest) (resp *order.SubmitOrderResponse, err error) {
	resp, err = s.OrderSqlManage.SubmitOrder(request.UserId, request.Items, request.Address, request.PaymentMethod)
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	return
}

// PayOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PayOrder(ctx context.Context, request *order.PayOrderRequest) (resp *order.PayOrderResponse, err error) {
	resp, err = s.OrderSqlManage.PayOrder(request.OrderId, request.PaymentMethod, request.PaymentDetails)
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	return
}

// CancelOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CancelOrder(ctx context.Context, request *order.CancelOrderRequest) (resp *order.CancelOrderResponse, err error) {
	resp, err = s.OrderSqlManage.CancelOrder(request.OrderId)
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	return
}

// GetOrders implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrders(ctx context.Context, request *order.GetOrdersRequest) (resp *order.GetOrdersResponse, err error) {
	resp, err = s.OrderSqlManage.GetOrders(request.UserId)
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	return
}

// GetOrderById implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrderById(ctx context.Context, request *order.GetOrderByIdRequest) (resp *order.GetOrderByIdResponse, err error) {
	resp, err = s.OrderSqlManage.GetOrderById(request.OrderId)
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	return
}

// CreateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (resp *order.CreateOrderResponse, err error) {
	// TODO: Your code here...
	return
}

// ListOrders implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrders(ctx context.Context, req *order.ListOrdersRequest) (resp *order.ListOrdersResponse, err error) {
	// TODO: Your code here...
	return
}

// GetOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrder(ctx context.Context, req *order.GetOrderRequest) (resp *order.GetOrderResponse, err error) {
	// TODO: Your code here...
	return
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidRequest) (resp *order.MarkOrderPaidResponse, err error) {
	// TODO: Your code here...
	return
}
