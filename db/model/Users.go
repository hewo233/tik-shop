package model

import (
	"time"
)

type Users struct {
	Id             int64      `gorm:"primaryKey;autoIncrement;column:Id" thrift:"id,1" frugal:"1,default,i64" json:"id"`
	Username       string     `gorm:"column:Username;size:255" thrift:"username,2" frugal:"2,default,string" json:"username"`
	HashedPassword string     `gorm:"column:HashedPassword;size:61"`
	Email          string     `gorm:"column:Email;size:255" thrift:"email,3" frugal:"3,default,string" json:"email"`
	Role           string     `gorm:"column:Role;size:50" thrift:"role,4" frugal:"4,default,string" json:"role"`
	CreatedAt      time.Time  `gorm:"column:CreatedAt;autoCreateTime" thrift:"createdAt,5" frugal:"5,default,string" json:"createdAt"`
	CartItems      []CartItem `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Orders         []Order    `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
