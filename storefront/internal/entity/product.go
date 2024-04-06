package entity

import "github.com/google/uuid"

type Product struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
}
