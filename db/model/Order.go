package model

import "time"

// Order Status
const (
	OrderStatusPending   int8 = 0 // 待支付
	OrderStatusPaid      int8 = 1 // 已支付/待发货
	OrderStatusShipped   int8 = 2 // 已发货
	OrderStatusCompleted int8 = 3 // 已完成
	OrderStatusCancelled int8 = 4 // 已取消
	OrderStatusRefunded  int8 = 5 // 已退款
)

type OrderAddress struct {
	CustomerName string `json:"customer_name"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
}

// Order 订单主表
type Order struct {
	ID         int64 `gorm:"primaryKey;column:id" json:"id"`
	CustomerID int64 `gorm:"column:customer_id;index;not null" json:"customer_id"`

	// 订单状态：0:待支付, 1:已支付/待发货, 2:已发货, 3:已完成, 4:已取消
	Status int8 `gorm:"column:status;default:0;index" json:"status"`
	// 订单总金额 (所有 Item 之和)
	TotalAmount int64 `gorm:"column:total_amount;type:bigint;not null" json:"total_amount"`

	// 收货信息快照 (不要直接关联 Address 表，因为用户修改地址不应影响历史订单)
	Address OrderAddress `gorm:"column:address;serializer:json" json:"address"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	Customer   *Customer   `gorm:"foreignKey:CustomerID;references:UserID" json:"-"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID;references:ID" json:"items"`
}

// OrderItem 订单明细表 (记录订单里买了什么，当时多少钱)
type OrderItem struct {
	ID         int64 `gorm:"primaryKey;column:id" json:"id"`
	OrderID    int64 `gorm:"column:order_id;index;not null" json:"order_id"`
	ProductID  int64 `gorm:"column:product_id;index;not null" json:"product_id"`
	MerchantID int64 `gorm:"column:merchant_id;index;not null" json:"merchant_id"`

	ProductName string `gorm:"column:product_name;size:255" json:"product_name"`
	Cost        int64  `gorm:"column:price;type:bigint;not null" json:"cost"` // 下单时的单价 (快照)
	Quantity    int64  `gorm:"column:quantity;not null" json:"quantity"`
	TotalCost   int64  `gorm:"column:total_cost;type:bigint;not null" json:"total_cost"` // 下单时的总价 (快照)

	// 关联
	Order *Order `gorm:"foreignKey:OrderID;references:ID" json:"-"`
}
