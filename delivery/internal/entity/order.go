package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	UUID      uuid.UUID   `json:"uuid"`
	Status    OrderStatus `json:"order_status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
