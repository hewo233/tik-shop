package model

type PaymentDetails struct {
	Id         int64  `gorm:"primaryKey;autoIncrement;column:Id" thrift:"id,1" frugal:"1,default,i64" json:"id"`
	OrderId    uint   `gorm:"column:OrderId"`
	CardNumber string `gorm:"column:CardNumber;size:20;not null" thrift:"cardNumber,1" frugal:"1,default,string" json:"cardNumber"`
	ExpiryDate string `gorm:"column:ExpiryDate;size:5;not null" thrift:"expiryDate,2" frugal:"2,default,string" json:"expiryDate"`
	Cvv        string `gorm:"column:Cvv;size:4;not null" thrift:"cvv,3" frugal:"3,default,string" json:"cvv"`
}
