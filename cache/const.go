package cache

import "time"

// Cache Key

const (
	ProductInfoKey  = "product:info:"  // + productID
	ProductStockKey = "product:stock:" // + productID
	ProductListKey  = "product:list:"  // + merchantID
)

const (
	NullPlaceholder = "*"
	ShortExpire     = 5 * time.Minute // 空值缓存时间
	DefaultExpire   = 24 * time.Hour  // 默认缓存时间
	ListExpire      = 1 * time.Hour   // 列表缓存时间
)
