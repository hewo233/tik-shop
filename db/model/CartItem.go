package model

type CartItem struct {
	Id        uint  `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductId int64 `gorm:"column:ProductId" json:"ProductId"`
	UserId    uint  `gorm:"column:UserId"` // 关联的用户 ID
	Quantity  int64 `gorm:"column:Quantity;not null" json:"Quantity"`
}
