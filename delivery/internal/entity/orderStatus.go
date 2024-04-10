package entity

type OrderStatus string

const (
	StatusCreated  OrderStatus = "created"
	StatusPrepare  OrderStatus = "prepare"
	StatusDelivery OrderStatus = "delivery"
	StatusAwait    OrderStatus = "await"
	StatusReceived OrderStatus = "received"
)
