package main

import (
	"context"
	"github.com/hewo/tik-shop/db/model"
	order "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/order"
	"github.com/jinzhu/copier"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct {
	OrderSqlManage
}
type OrderSqlManage interface {
	CreateOrder(order *order.CreateOrderRequest) (int64, error)
	GetOrder(orderID int64, customerID int64) (*model.Order, error)
	ListOrders(customerID int64, offset int, limit int, status int) ([]*model.Order, error)
	CancelOrder(orderID int64, customerID int64) error
}

// CreateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (resp *order.CreateOrderResponse, err error) {
	id, err := s.OrderSqlManage.CreateOrder(req)
	if err != nil {
		return nil, err
	}
	resp = &order.CreateOrderResponse{
		OrderId: id,
	}
	return resp, nil
}

// ListOrders implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrders(ctx context.Context, req *order.ListOrdersRequest) (resp *order.ListOrdersResponse, err error) {
	offset := (req.Page - 1) * req.PageSize
	limit := req.PageSize

	ors, err := s.OrderSqlManage.ListOrders(req.CustomerId, int(offset), int(limit), int(*req.Status))
	if err != nil {
		return nil, err
	}
	resp = &order.ListOrdersResponse{}
	for _, ord := range ors {
		o := &order.Order{}
		if err := copier.Copy(o, ord); err != nil {
			return nil, err
		}
		resp.Orders = append(resp.Orders, o)
		resp.Total++
	}
	return resp, nil
}

// GetOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrder(ctx context.Context, req *order.GetOrderRequest) (resp *order.GetOrderResponse, err error) {
	ord, err := s.OrderSqlManage.GetOrder(req.OrderId, req.CustomerId)
	if err != nil {
		return nil, err
	}
	resp = &order.GetOrderResponse{}
	resp.Order = &order.Order{}
	if err := copier.Copy(resp.Order, ord); err != nil {
		return nil, err
	}
	return resp, nil
}

// CancelOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CancelOrder(ctx context.Context, req *order.CancelOrderRequest) (resp *order.CancelOrderResponse, err error) {
	err = s.OrderSqlManage.CancelOrder(req.OrderId, req.CustomerId)
	if err != nil {
		return nil, err
	}
	resp = &order.CancelOrderResponse{
		Success: true,
	}
	return resp, nil
}
