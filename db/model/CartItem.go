package model

type CartItem struct {
	ProductId int64 `gorm:"primaryKey;column:ProductId" json:"ProductId"`
	Quantity  int64 `gorm:"column:Quantity" json:"Quantity"`
}
