package model

import "time"

type OrderStatus int64

const (
	OrderStatus_PENDING   OrderStatus = 0
	OrderStatus_PAID      OrderStatus = 1
	OrderStatus_CANCELLED OrderStatus = 2
)

// OrderItem 结构体映射数据库表
type OrderItem struct {
	Id        int64   `gorm:"primaryKey;autoIncrement;column:Id" thrift:"id,1" frugal:"1,default,i64" json:"id"`
	ProductId int64   `gorm:"column:ProductId;not null" thrift:"productId,1" frugal:"1,default,i64" json:"productId"`
	Quantity  int64   `gorm:"column:Quantity;not null" thrift:"quantity,2" frugal:"2,default,i64" json:"quantity"`
	Price     float64 `gorm:"column:Price;not null" thrift:"price,3" frugal:"3,default,double" json:"price"`
	OrderId   int64   `gorm:"column:OrderId;not null;" json:"orderId"` // 外键引用 Order
}

// Order 结构体映射数据库表
type Order struct {
	Id             int64          `gorm:"primaryKey;autoIncrement;column:Id" thrift:"orderId,1" frugal:"1,default,i64" json:"orderId"`
	UserId         int64          `gorm:"column:UserId;not null"`
	Status         OrderStatus    `gorm:"column:Status;not null" thrift:"status,2" frugal:"2,default,OrderStatus" json:"status"`
	TotalAmount    float64        `gorm:"column:TotalAmount;not null" thrift:"totalAmount,3" frugal:"3,default,double" json:"totalAmount"`
	CreatedAt      time.Time      `gorm:"column:CreatedAt;autoCreateTime" thrift:"createdAt,4" frugal:"4,default,string" json:"createdAt"`
	PaymentMethod  string         `gorm:"column:PaymentMethod;not null"`
	Items          []OrderItem    `gorm:"foreignKey:OrderId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" thrift:"items,5" frugal:"5,default,list<OrderItem>" json:"items"`
	Address        Address        `gorm:"foreignKey:OrderId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PaymentDetails PaymentDetails `gorm:"foreignKey:OrderId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
