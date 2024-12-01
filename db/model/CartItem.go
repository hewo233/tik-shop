package model

type CartItem struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID int64   `gorm:"column:ProductID" json:"ProductId"`
	UserID    uint    `gorm:"column:UserID"` // 关联的用户 ID
	Quantity  int64   `gorm:"column:Quantity" json:"Quantity"`
	Users     Users   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`     // 关联用户
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnDelete:SET NULL;"` // 关联商品
}
