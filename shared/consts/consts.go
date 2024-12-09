package consts

import "time"

const (
	CorsAddress = "localhost:8080"

	AccountId = "accountID"

	Issuer         = "tik-shop"
	User           = "user"
	ID             = "id"
	Admin          = "admin"
	SevenDays      = time.Hour * 24 * 7
	ThirtyDays     = time.Hour * 24 * 30
	ApiConfigPath  = "./config.yaml"
	UserConfigPath = "./rpc/user-service/config.yaml"
)
