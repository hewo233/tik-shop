package model

import "time"

type CartItem struct {
	ID int64 `gorm:"primaryKey;column:id" json:"id"`

	// 归属人
	CustomerID int64 `gorm:"column:customer_id;index;not null" json:"customer_id"`

	// 买了哪个商品
	ProductID int64 `gorm:"column:product_id;index;not null" json:"product_id"`

	// 冗余 MerchantID：方便前端按“店铺”分组显示购物车，减少连表查询
	MerchantID int64 `gorm:"column:merchant_id;index;not null" json:"merchant_id"`

	// 数量
	Quantity int `gorm:"column:quantity;default:1" json:"quantity"`

	// 选中状态：很多电商允许购物车里有东西，但下单时只勾选一部分
	Selected bool `gorm:"column:selected;default:true" json:"selected"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	// 关联
	Customer *Customer `gorm:"foreignKey:CustomerID;references:UserID" json:"-"`
	Product  *Product  `gorm:"foreignKey:ProductID;references:ID" json:"product"` // 这里需要 Preload 出来，因为要显示当前图片和标题
}
