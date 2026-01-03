package main

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/hewo/tik-shop/db/model"
	order "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/order"
	"github.com/hewo/tik-shop/mq"
	"github.com/jinzhu/copier"
	"log"
	"time"
)

var idNode *snowflake.Node

func init() {
	var err error
	idNode, err = snowflake.NewNode(1) // 节点 ID 1
	if err != nil {
		panic(err)
	}
}

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct {
	OrderSqlManage
	ProductCacheManage
}
type OrderSqlManage interface {
	CreateOrder(order *order.CreateOrderRequest) (int64, error)
	GetOrder(orderID int64, customerID int64) (*model.Order, error)
	ListOrders(customerID int64, offset int, limit int, status int) ([]*model.Order, error)
	CancelOrder(orderID int64, customerID int64) error
}
type ProductCacheManage interface {
	DeductStock(ctx context.Context, productIDs []int64, quantities []int64) error
}

// CreateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (resp *order.CreateOrderResponse, err error) {
	orderID := idNode.Generate().Int64()

	items := make([]mq.OrderItem, len(req.Items))

	ProductIDs := make([]int64, len(req.Items))
	Quantities := make([]int64, len(req.Items))
	for i, item := range req.Items {
		ProductIDs[i] = item.ProductId
		Quantities[i] = item.Quantity
		items[i] = mq.OrderItem{
			ProductId: item.ProductId,
			Count:     item.Quantity,
		}
	}

	err = s.ProductCacheManage.DeductStock(ctx, ProductIDs, Quantities)
	if err != nil {
		return nil, err
	}

	// send message to mq
	msg := mq.OrderMessage{
		OrderId:   orderID,
		Uid:       req.CustomerId,
		Items:     items,
		Timestamp: time.Now().Unix(),
	}

	if err := mq.Producer.SendOrderMsg(ctx, msg); err != nil {
		log.Println("Failed to send order message to MQ:", err)
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
