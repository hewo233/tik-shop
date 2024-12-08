package consts

import "time"

const (
	CorsAddress = "localhost:8080"

	AccountId = "accountID"

	Issuer         = "tik-shop"
	User           = "user"
	ID             = "id"
	SevenDays      = time.Hour * 24 * 7
	ApiConfigPath  = "./config.yaml"
	UserConfigPath = "./rpc/user-service/config.yaml"
)
