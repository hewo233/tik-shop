package superquery

import (
	"errors"
	"github.com/hewo/tik-shop/db/model"
	"github.com/hewo/tik-shop/db/query"
	"github.com/hewo/tik-shop/mq"
)

type MqSqlManageImpl struct{}

func (m *MqSqlManageImpl) MqCreateOrder(orderId int64, uid int64, items []mq.OrderItem) error {
	// create order in db

	productIds := make([]int64, 0, len(items))
	for _, item := range items {
		productIds = append(productIds, item.ProductId)
	}

	err := query.Q.Transaction(func(tx *query.Query) error {

		u := tx.User

		user, err := u.Preload(u.Customer).Where(u.ID.Eq(uid)).First()
		if err != nil {
			return err
		}
		p := tx.Product
		products, err := p.Where(p.ID.In(productIds...)).Find()
		if err != nil {
			return err
		}
		productMap := make(map[int64]*model.Product)
		for _, product := range products {
			productMap[product.ID] = product
		}

		var orderItems []model.OrderItem
		var totalAmount int64 = 0

		for _, reqItem := range items {
			prod, ok := productMap[reqItem.ProductId]
			if !ok {
				return errors.New("product id not found")
			}
			info, err := p.Where(p.ID.Eq(reqItem.ProductId), p.Stock.Gte(reqItem.Count)).UpdateSimple(p.Stock.Sub(reqItem.Count))
			if err != nil {
				return err
			}
			if info.RowsAffected == 0 {
				return errors.New("product stock not enough")
			}

			orderItem := model.OrderItem{
				ProductID:   prod.ID,
				MerchantID:  prod.MerchantID,
				Quantity:    reqItem.Count,
				ProductName: prod.Name,
				Cost:        prod.Price,
				TotalCost:   prod.Price * reqItem.Count,
			}
			totalAmount += orderItem.TotalCost
			orderItems = append(orderItems, orderItem)
		}

		order := model.Order{
			ID:          orderId,
			CustomerID:  uid,
			Status:      model.OrderStatusWaiting,
			TotalAmount: 0,
			Address: model.OrderAddress{
				CustomerName: user.Username,
				Phone:        user.Customer.Phone,
				Address:      user.Customer.Address,
			},
			OrderItems: orderItems,
		}
		if err := tx.Order.Create(&order); err != nil {
			return err
		}

		return nil
	})
	return err
}
