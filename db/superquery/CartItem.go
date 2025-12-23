package superquery

import (
	"errors"
	"fmt"
	"github.com/hewo/tik-shop/db/model"
	"github.com/hewo/tik-shop/db/query"
	"gorm.io/gorm"
)

var c = &query.Q.CartItem

type CartSqlManageImpl struct{}

func NewCartItemSqlManageImpl() *CartSqlManageImpl {
	return &CartSqlManageImpl{}
}

func (m *CartSqlManageImpl) GetCart(customerID int64) ([]*model.CartItem, error) {
	items, err := c.Preload(c.Product).Where(c.CustomerID.Eq(customerID)).Order(c.CreatedAt.Desc()).Find()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (m *CartSqlManageImpl) AddToCart(customerID int64, productID int64, quantity int64) (int64, error) {
	p := query.Q.Product

	product, err := p.Where(p.ID.Eq(productID)).First()
	if err != nil {
		return -1, err
	}

	// 2. 尝试查询购物车现有记录
	item, err := c.Where(c.CustomerID.Eq(customerID), c.ProductID.Eq(productID)).First()

	// Case A: 购物车中不存在 -> 新增
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 考虑到购物车并没有那么严重并发问题，这里就不搞分布式锁
		if quantity > product.Stock {
			return -1, fmt.Errorf("product stock is not enough")
		}
		newItem := &model.CartItem{
			CustomerID: customerID,
			ProductID:  productID,
			MerchantID: product.MerchantID,
			Quantity:   quantity,
			Selected:   false,
		}
		err2 := c.Create(newItem)
		if err2 != nil {
			return -1, err2
		}
		return newItem.ID, nil
	} else if err != nil {
		return -1, err
	}

	info, err := c.Where(
		c.ID.Eq(item.ID),
		c.Quantity.Add(quantity).Lte(product.Stock),
	).Updates(map[string]interface{}{
		"quantity": gorm.Expr("quantity + ?", quantity),
		"selected": true,
	})

	if err != nil {
		return -1, err
	}

	// 如果受影响行数为 0，说明 WHERE 条件不满足（即库存不足）
	if info.RowsAffected == 0 {
		return -1, fmt.Errorf("库存不足或商品状态变化")
	}

	return item.ID, nil
}

func (m *CartSqlManageImpl) UpdateQuantity(cartItemID int64, customerID int64, newQuantity int64) error {
	if newQuantity <= 0 {
		// delete
		info, err := c.Where(c.ID.Eq(cartItemID), c.CustomerID.Eq(customerID)).Delete()
		if err != nil {
			return err
		}
		if info.RowsAffected == 0 {
			return fmt.Errorf("cart item not found")
		}
	}
	// update
	info, err := c.Where(c.ID.Eq(cartItemID), c.CustomerID.Eq(customerID)).Update(c.Quantity, newQuantity)
	if err != nil {
		return err
	}
	if info.RowsAffected == 0 {
		return fmt.Errorf("cart item not found")
	}
	return nil
}

func (m *CartSqlManageImpl) ToggleSelect(cartItemIDs []int64, customerID int64, selected bool) error {
	if len(cartItemIDs) == 0 {
		return nil
	}

	info, err := c.Where(c.CustomerID.Eq(customerID), c.ID.In(cartItemIDs...)).Update(c.Selected, selected)
	if err != nil {
		return err
	}
	if info.RowsAffected == 0 {
		return fmt.Errorf("cart item not found or no permission")
	}
	return nil
}

func (m *CartSqlManageImpl) RemoveItems(cartItemIDs []int64, customerID int64) error {
	if len(cartItemIDs) == 0 {
		return nil
	}

	info, err := c.Where(c.CustomerID.Eq(customerID), c.ID.In(cartItemIDs...)).Delete()
	if err != nil {
		return err
	}
	if info.RowsAffected == 0 {
		return fmt.Errorf("cart item not found or no permission")
	}
	return nil
}

func (m *CartSqlManageImpl) ClearCart(customerID int64) error {
	info, err := c.Where(c.CustomerID.Eq(customerID)).Delete()
	if err != nil {
		return err
	}
	if info.RowsAffected == 0 {
		return fmt.Errorf("cart is already empty")
	}
	return nil
}
