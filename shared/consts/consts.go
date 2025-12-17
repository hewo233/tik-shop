package consts

import "time"

const (
	CorsAddress = "localhost:8080"

	AccountID = "accountID"

	Issuer         = "tik-shop"
	User           = "user"
	ID             = "id"
	Admin          = "admin"
	SevenDays      = time.Hour * 24 * 7
	ThirtyDays     = time.Hour * 24 * 30

	RoleCustomer  = "customer"
	RoleMerchant  = "merchant"
	RoleAdmin     = "admin"
)
