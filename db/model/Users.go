package model

import (
	"time"
)

// User 基础用户表
type User struct {
	ID             int64     `gorm:"primaryKey;column:id" json:"id"`
	Username       string    `gorm:"column:username;size:255;not null;unique" json:"username"`
	HashedPassword string    `gorm:"column:hashed_password;size:255;not null"`
	Email          string    `gorm:"column:email;size:255;not null;uniqueIndex" json:"email"`
	Role           string    `gorm:"column:role;size:50;not null" json:"role"`    // ex: "customer", "merchant", "admin"
	Status         int8      `gorm:"column:status;default:1;index" json:"status"` // 1:ok, 0:deleted, 2:banned
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	// 1-to-1 ：User 是源头
	Customer *Customer `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"customer,omitempty"`
	Merchant *Merchant `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"merchant,omitempty"`
	Admin    *Admin    `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"admin,omitempty"`
}

// Customer 消费者扩展表
type Customer struct {
	UserID  int64  `gorm:"primaryKey;column:user_id" json:"user_id"`
	Address string `gorm:"column:address;size:512" json:"address"`
	Phone   string `gorm:"column:phone;size:20" json:"phone"`

	User   *User      `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Orders []Order    `gorm:"foreignKey:CustomerID;references:UserID" json:"orders,omitempty"`
	Cart   []CartItem `gorm:"foreignKey:CustomerID;references:UserID" json:"cart,omitempty"`
}

// Merchant 商家扩展表
type Merchant struct {
	UserID   int64  `gorm:"primaryKey;column:user_id" json:"user_id"`
	Address  string `gorm:"column:address;size:512" json:"address"`
	ShopName string `gorm:"column:shop_name;size:255" json:"shop_name"`

	User     *User     `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Products []Product `gorm:"foreignKey:MerchantID;references:UserID" json:"products,omitempty"`
}

// Admin 管理员扩展表
type Admin struct {
	UserID int64 `gorm:"primaryKey;column:user_id" json:"user_id"`
	Level  int   `gorm:"column:level;default:1" json:"level"`

	User *User `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
}
