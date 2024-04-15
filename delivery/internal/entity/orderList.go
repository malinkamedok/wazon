package entity

import "github.com/google/uuid"

type OrderList struct {
	UUID   uuid.UUID   `json:"uuid"`
	Status OrderStatus `json:"order_status"`
}
