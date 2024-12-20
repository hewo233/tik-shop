package superquery

import (
	"errors"
	"fmt"

	"github.com/hewo/tik-shop/db/model"
	"github.com/hewo/tik-shop/db/query"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/order"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var o = &query.Q.Order

type OrderSqlManageImpl struct{}

func NewOrderSqlManageImpl() *OrderSqlManageImpl {
	return &OrderSqlManageImpl{}
}
func (m *OrderSqlManageImpl) SubmitOrder(UserId int64, Items []*order.OrderItem, Address *order.Address, PaymentMethod string) (*order.SubmitOrderResponse, error) {

	// 计算总金额，这里假设每个商品的价格存储在 Items 中
	var totalAmount float64
	for _, item := range Items {
		// 假设每个 item 有 Price 字段
		totalAmount += item.Price * float64(item.Quantity)
	}
	ADDRESS := model.Address{OrderId: 0}
	err := copier.Copy(&ADDRESS, &Address)
	if err != nil {
		return nil, fmt.Errorf("copier.Copy Address: %w", err)
	}

	ORDER := model.Order{
		UserId:        UserId,
		Status:        model.OrderStatus_PENDING, // 假设默认状态为 "Pending"
		TotalAmount:   totalAmount,
		PaymentMethod: PaymentMethod,
		Items:         make([]model.OrderItem, len(Items)),
		Address:       ADDRESS, // 假设 Address 已经是一个完整的 Address 对象
	}
	ADDRESS.OrderId = uint(ORDER.Id)

	// 转换 Items 数据到 OrderItem 结构体
	for i, item := range Items {
		ORDER.Items[i].OrderId = ORDER.Id
		err = copier.Copy(&ORDER.Items[i], &item)
		if err != nil {
			return nil, fmt.Errorf("copier.Copy OrderItems: %w", err)
		}
	}
	err = o.Create(&ORDER)
	if err != nil {
		return nil, fmt.Errorf("o.Create: %w", err)
	}
	return &order.SubmitOrderResponse{
		OrderId: ORDER.Id,
		Message: "Submit order successfully",
	}, nil
}

func (m *OrderSqlManageImpl) PayOrder(orderId int64, PaymentMethod string, PaymentDetails *order.PaymentDetails) (*order.PayOrderResponse, error) {
	//查找有没有相应的orderId
	_, err := o.Where(o.Id.Eq(orderId)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("order not found")
		}
	}
	//确定订单处于未支付状态
	_, err = o.Where(o.Id.Eq(orderId), o.Status.Eq(0)).First()
	if err != nil {
		return nil, fmt.Errorf("order cannot be paid, current status: %v", o.Status)
	}

	_, err = o.Where(o.PaymentMethod.Eq(PaymentMethod), o.Id.Eq(orderId)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("paymethod dont't match")
		}
	}
	//更新支付状态
	_, err = o.Update(o.Status, 1)
	if err != nil {
		return nil, fmt.Errorf("update status fail")
	}

	// 记录支付详情（如果有的话）
	if PaymentDetails != nil {
		paymentDetails := model.PaymentDetails{OrderId: uint(orderId)}

		err = copier.Copy(&paymentDetails, &PaymentDetails)
		if err != nil {
			return nil, fmt.Errorf("copier.Copy PaymentDetails: %w", err)
		}

		err = query.Q.PaymentDetails.Create(&paymentDetails)
		if err != nil {
			return nil, fmt.Errorf("create payment details fail")
		}
	}
	return &order.PayOrderResponse{
		Message: "pay order successfully",
	}, nil
}

func (m *OrderSqlManageImpl) CancelOrder(orderId int64) (*order.CancelOrderResponse, error) {
	// 1. 查找数据库中的订单
	if _, err := o.Where(o.Id.Eq(orderId)).First(); err != nil {
		// 如果未找到订单，返回错误
		return nil, fmt.Errorf("order not found: %v", err)
	}

	// 2. 检查订单状态是否为可取消（假设只有 "PENDING" 状态可以取消）
	_, err := o.Where(o.Id.Eq(orderId), o.Status.Eq(0)).First()
	if err != nil {
		return nil, fmt.Errorf("order status is not cancelable")
	}

	// 3. 更新订单状态为已取消
	_, err = o.Where(o.Id.Eq(orderId)).Delete()
	if err != nil {
		// 如果更新失败，返回错误
		return nil, fmt.Errorf("failed to cancel order: %v", err)
	}
	return &order.CancelOrderResponse{
		Message: "cancel order successfully",
	}, nil
}

func (m *OrderSqlManageImpl) GetOrders(userId int64) (*order.GetOrdersResponse, error) {
	ORDERS, err := o.Where(o.UserId.Eq(userId)).Find()
	if err != nil {
		return nil, fmt.Errorf("failed to find orders: %v", err)
	}
	var orders = make([]*order.Order, len(ORDERS))
	for i, ORDER := range ORDERS {
		orders[i] = &order.Order{
			Status:    order.OrderStatus(ORDER.Status),
			CreatedAt: ORDER.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}
	err = copier.CopyWithOption(&orders, ORDERS, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, fmt.Errorf("copier.Copy Orders: %w", err)
	}

	return &order.GetOrdersResponse{
		Orders: orders, // 返回查询到的订单
	}, nil
}

func (m *OrderSqlManageImpl) GetOrderById(orderId int64) (*order.GetOrderByIdResponse, error) {
	ORDER, err := o.Where(o.Id.Eq(orderId)).First()
	if err != nil {
		return nil, fmt.Errorf("failed to find order: %v", err)
	}
	orderById := &order.Order{
		Status:    order.OrderStatus(ORDER.Status),
		CreatedAt: ORDER.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	err = copier.CopyWithOption(&orderById, ORDER, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, fmt.Errorf("copier.Copy OrderById: %w", err)
	}
	return &order.GetOrderByIdResponse{
		Order: orderById,
	}, nil
}
