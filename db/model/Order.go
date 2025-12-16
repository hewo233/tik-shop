package model

import "time"

// Order Status
const (
	OrderStatusPending   int8 = 0 // 待支付
	OrderStatusPaid      int8 = 1 // 已支付/待发货
	OrderStatusShipped   int8 = 2 // 已发货
	OrderStatusCompleted int8 = 3 // 已完成
	OrderStatusCancelled int8 = 4 // 已取消
)

// Order 订单主表
type Order struct {
	ID         int64 `gorm:"primaryKey;column:id" json:"id"`
	CustomerID int64 `gorm:"column:customer_id;index;not null" json:"customer_id"`

	// 订单号，通常是一个唯一字符串，不仅限于数字 ID
	OrderNo string `gorm:"column:order_no;size:64;unique;not null" json:"order_no"`

	// 订单总金额 (所有 Item 之和)
	TotalAmount int64 `gorm:"column:total_amount;type:decimal;not null" json:"total_amount"`

	// 订单状态：0:待支付, 1:已支付/待发货, 2:已发货, 3:已完成, 4:已取消
	Status int8 `gorm:"column:status;default:0;index" json:"status"`

	// 收货信息快照 (不要直接关联 Address 表，因为用户修改地址不应影响历史订单)
	SnapshotAddress string `gorm:"column:snapshot_address;size:1024" json:"snapshot_address"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	// 关联
	Customer   *Customer   `gorm:"foreignKey:CustomerID;references:UserID" json:"-"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID;references:ID" json:"items"`
}

// OrderItem 订单明细表 (记录订单里买了什么，当时多少钱)
type OrderItem struct {
	ID         int64 `gorm:"primaryKey;column:id" json:"id"`
	OrderID    int64 `gorm:"column:order_id;index;not null" json:"order_id"`
	ProductID  int64 `gorm:"column:product_id;index;not null" json:"product_id"`
	MerchantID int64 `gorm:"column:merchant_id;index;not null" json:"merchant_id"` // 冗余字段，方便拆单或商家查询

	ProductName string  `gorm:"column:product_name;size:255" json:"product_name"`      // 冗余字段，防止商品被删后订单显示空白
	Price       float64 `gorm:"column:price;type:decimal(10,2);not null" json:"price"` // 下单时的单价 (快照)
	Quantity    int     `gorm:"column:quantity;not null" json:"quantity"`

	// 关联
	Order   *Order   `gorm:"foreignKey:OrderID;references:ID" json:"-"`
	Product *Product `gorm:"foreignKey:ProductID;references:ID" json:"-"`
}
