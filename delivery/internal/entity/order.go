package entity

import "github.com/google/uuid"

type Order struct {
	UUID      uuid.UUID   `json:"uuid"`
	Status    OrderStatus `json:"order_status"`
	CreatedAt uuid.Time   `json:"created_at"`
	UpdatedAt uuid.Time   `json:"updated_at"`
}
