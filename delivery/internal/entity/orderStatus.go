package entity

import "fmt"

type OrderStatus string

const (
	StatusCreated  OrderStatus = "created"
	StatusPrepare  OrderStatus = "prepare"
	StatusDelivery OrderStatus = "delivery"
	StatusAwait    OrderStatus = "await"
	StatusReceived OrderStatus = "received"
)

func StringToStatus(in string) (OrderStatus, error) {
	switch in {
	case "created":
		return StatusCreated, nil
	case "prepare":
		return StatusPrepare, nil
	case "delivery":
		return StatusDelivery, nil
	case "await":
		return StatusAwait, nil
	case "received":
		return StatusReceived, nil
	default:
		return OrderStatus(in), fmt.Errorf("Unknown status %s", in)
	}
}
