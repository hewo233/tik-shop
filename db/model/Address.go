package model

type Address struct {
	Id         int64  `gorm:"primaryKey;autoIncrement;column:Id" thrift:"id,1" frugal:"1,default,i64" json:"id"`
	OrderId    uint   `gorm:"column:OrderId"`
	Street     string `gorm:"column:Street;not null;size:255" thrift:"street,1" frugal:"1,default,string" json:"street"`
	City       string `gorm:"column:City;not null;size:255" thrift:"city,2" frugal:"2,default,string" json:"city"`
	PostalCode string `gorm:"column:PostalCode;not null;size:20" thrift:"postalCode,3" frugal:"3,default,string" json:"postalCode"`
	Country    string `gorm:"column:Country;not null;size:100" thrift:"country,4" frugal:"4,default,string" json:"country"`
}
