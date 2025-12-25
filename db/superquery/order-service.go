package superquery

import (
	"fmt"
	"github.com/hewo/tik-shop/db/model"
	"github.com/hewo/tik-shop/db/query"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/order"
)

var o = &query.Q.Order

type OrderSqlManageImpl struct{}

func NewOrderSqlManageImpl() *OrderSqlManageImpl {
	return &OrderSqlManageImpl{}
}

func (m *OrderSqlManageImpl) CreateOrder(order *order.CreateOrderRequest) (int64, error) {

	var newOrderID int64

	err := query.Q.Transaction(func(tx *query.Query) error {

		// get user address
		u := tx.User
		user, err := u.Preload(u.Customer).Where(u.ID.Eq(order.CustomerId)).First()
		if err != nil {
			return err
		}

		productIDs := make([]int64, 0, len(order.Items))
		for _, item := range order.Items {
			productIDs = append(productIDs, item.ProductId)
		}

		p := tx.Product
		products, err := p.Where(p.ID.In(productIDs...)).Find()
		if err != nil {
			return err
		}
		productMap := make(map[int64]*model.Product)
		for _, product := range products {
			productMap[product.ID] = product
		}

		var orderItems []model.OrderItem
		var totalAmount int64 = 0
		for _, reqItem := range order.Items {
			prod, ok := productMap[reqItem.ProductId]
			if !ok {
				return fmt.Errorf("product id %d not found", reqItem.ProductId)
			}

			info, err := p.Where(p.ID.Eq(reqItem.ProductId), p.Stock.Gte(reqItem.Quantity)).UpdateSimple(p.Stock.Sub(reqItem.Quantity))
			if err != nil {
				return err
			}
			if info.RowsAffected == 0 {
				return fmt.Errorf("product id %d not enough", reqItem.ProductId)
			}

			orderItems = append(orderItems, model.OrderItem{
				ProductID:   prod.ID,
				MerchantID:  prod.MerchantID,
				Quantity:    reqItem.Quantity,
				ProductName: prod.Name,
				Cost:        prod.Price,
				TotalCost:   prod.Price * reqItem.Quantity,
			})
			totalAmount += prod.Price * reqItem.Quantity
		}
		newOrder := &model.Order{
			CustomerID:  order.CustomerId,
			Status:      model.OrderStatusPending,
			TotalAmount: totalAmount,
			Address: model.OrderAddress{
				CustomerName: user.Username,
				Phone:        user.Customer.Phone,
				Address:      user.Customer.Address,
			},
			OrderItems: orderItems,
		}
		if err := tx.Order.Create(newOrder); err != nil {
			return err
		}
		newOrderID = newOrder.ID
		return nil
	})

	return newOrderID, err
}

func (m *OrderSqlManageImpl) GetOrder(orderID int64, customerID int64) (*model.Order, error) {
	ord, err := o.Preload(o.OrderItems).Where(o.ID.Eq(orderID), o.CustomerID.Eq(customerID)).First()
	return ord, err
}

func (m *OrderSqlManageImpl) ListOrders(customerID int64, offset int, limit int, status int) ([]*model.Order, error) {

	ords := make([]*model.Order, 0)
	var err error

	if status == -1 {
		ords, err = o.Preload(o.OrderItems).Where(o.CustomerID.Eq(customerID)).Offset(offset).Limit(limit).Find()
	} else {
		ords, err = o.Preload(o.OrderItems).Where(o.CustomerID.Eq(customerID), o.Status.Eq(int8(status))).Offset(offset).Limit(limit).Find()
	}
	return ords, err
}

func (m *OrderSqlManageImpl) CancelOrder(orderID int64, customerID int64) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		o := tx.Order
		p := tx.Product

		orderInfo, err := o.Preload(o.OrderItems).
			Where(o.ID.Eq(orderID), o.CustomerID.Eq(customerID), o.Status.Eq(model.OrderStatusPending)).
			First()

		if err != nil {
			return fmt.Errorf("order not found or not pending: %w", err)
		}

		_, err = o.Where(o.ID.Eq(orderID)).Update(o.Status, model.OrderStatusCancelled)
		if err != nil {
			return err
		}

		for _, item := range orderInfo.OrderItems {
			_, err := p.Where(p.ID.Eq(item.ProductID)).UpdateSimple(p.Stock.Add(item.Quantity))
			if err != nil {
				return err
			}
		}

		return nil
	})
}
