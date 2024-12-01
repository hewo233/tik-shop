package model

type Product struct {
	Id          int64      `gorm:"primaryKey;autoIncrement;column:Id" thrift:"id,1" frugal:"1,default,i64" json:"id"`
	Name        string     `gorm:"column:Name" thrift:"name,2" frugal:"2,default,string" json:"name"`
	Price       float64    `gorm:"column:Price" thrift:"price,3" frugal:"3,default,double" json:"price"`
	Stock       int64      `gorm:"column:Stock" thrift:"stock,4" frugal:"4,default,i64" json:"stock"`
	Description string     `gorm:"column:Description" thrift:"description,5" frugal:"5,default,string" json:"description"`
	CartItems   []CartItem `gorm:"foreignKey:ProductID"`
}
