package model

import "time"

type Product struct {
	ID          int64  `gorm:"primaryKey;column:id" json:"id"`
	MerchantID  int64  `gorm:"column:merchant_id;index;not null" json:"merchant_id"` // 外键关联 Merchant
	Name        string `gorm:"column:name;size:255;not null;index" json:"name"`
	Description string `gorm:"column:description;type:text" json:"description"`

	Price  int64 `gorm:"column:price;not null" json:"price"`
	Stock  int64 `gorm:"column:stock;default:0" json:"stock"`
	Status int8  `gorm:"column:status;default:1;index" json:"status"` // 1:上架, 0:删除 2:下架 3:售罄

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	// 关联
	Merchant *Merchant `gorm:"foreignKey:MerchantID;references:UserID" json:"-"` // json:"-" 避免循环嵌套输出
}
