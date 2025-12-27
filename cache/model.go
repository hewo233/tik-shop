package cache

type CachedProduct struct {
	ID          int64  `json:"id"`
	MerchantID  int64  `json:"merchant_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
	Status      int8   `json:"status"`
}
