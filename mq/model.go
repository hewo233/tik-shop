package mq

// OrderMessage 定义投递到 MQ 的消息体

type OrderItem struct {
	ProductId int64 `json:"product_id"`
	Count     int64    `json:"count"`
}
type OrderMessage struct {
	OrderId   int64       `json:"order_id"`
	Uid       int64       `json:"uid"`
	Items     []OrderItem `json:"items"`
	Timestamp int64       `json:"timestamp"`
}
