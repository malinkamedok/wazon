package entity

import "github.com/google/uuid"

type Product struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
}
