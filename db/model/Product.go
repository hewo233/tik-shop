package model

type Product struct {
	Id          int64   `gorm:"primaryKey;autoIncrement;column:Id" thrift:"id,1" frugal:"1,default,i64" json:"id"`
	Name        string  `gorm:"column:Name;not null" thrift:"name,2" frugal:"2,default,string" json:"name"`
	Price       float64 `gorm:"column:Price;not null" thrift:"price,3" frugal:"3,default,double" json:"price"`
	Stock       int64   `gorm:"column:Stock;not null" thrift:"stock,4" frugal:"4,default,i64" json:"stock"`
	Description string  `gorm:"column:Description;not null" thrift:"description,5" frugal:"5,default,string" json:"description"`

	CartItems []CartItem `gorm:"foreignKey:ProductId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
