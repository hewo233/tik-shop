package main

import (
	"context"
	"fmt"
	"time"

	order "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/order"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// SubmitOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) SubmitOrder(ctx context.Context, request *order.SubmitOrderRequest) (resp *order.SubmitOrderResponse, err error) {
	// TODO: Your code here...
	resp = &order.SubmitOrderResponse{
		OrderId: 12345, // Fake order ID
		Message: "Order submitted successfully",
	}
	return
}

// PayOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PayOrder(ctx context.Context, request *order.PayOrderRequest) (resp *order.PayOrderResponse, err error) {
	// TODO: Your code here...
	resp = &order.PayOrderResponse{
		Message: "Payment processed successfully for Order ID " + fmt.Sprint(request.OrderId),
	}
	return
}

// CancelOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CancelOrder(ctx context.Context, request *order.CancelOrderRequest) (resp *order.CancelOrderResponse, err error) {
	// TODO: Your code here...
	resp = &order.CancelOrderResponse{
		Message: "Order ID " + fmt.Sprint(request.OrderId) + " has been cancelled",
	}
	return
}

// GetOrders implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrders(ctx context.Context, request *order.GetOrdersRequest) (resp *order.GetOrdersResponse, err error) {
	// TODO: Your code here...
	resp = &order.GetOrdersResponse{
		Orders: []*order.Order{
			{
				OrderId:     1,
				Status:      order.OrderStatus_PAID,
				TotalAmount: 299.99,
				CreatedAt:   time.Now().Format(time.RFC3339),
				Items: []*order.OrderItem{
					{ProductId: 101, Quantity: 2, Price: 9999},
					{ProductId: 102, Quantity: 1, Price: 19999},
				},
			},
			{
				OrderId:     2,
				Status:      order.OrderStatus_CANCELLED,
				TotalAmount: 49.99,
				CreatedAt:   time.Now().Add(-24 * time.Hour).Format(time.RFC3339),
				Items: []*order.OrderItem{
					{ProductId: 103, Quantity: 1, Price: 4999},
				},
			},
		},
	}
	return
}

// GetOrderById implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrderById(ctx context.Context, request *order.GetOrderByIdRequest) (resp *order.GetOrderByIdResponse, err error) {
	// TODO: Your code here...
	resp = &order.GetOrderByIdResponse{
		Order: &order.Order{
			OrderId:     request.OrderId,
			Status:      order.OrderStatus_PAID,
			TotalAmount: 149.99,
			CreatedAt:   time.Now().Add(-48 * time.Hour).Format(time.RFC3339),
			Items: []*order.OrderItem{
				{ProductId: 104, Quantity: 3, Price: 4999},
			},
		},
	}
	return
}
