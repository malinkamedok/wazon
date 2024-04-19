package entity

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	UUID      uuid.UUID `json:"uuid"`
	Status    string    `json:"order_status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
