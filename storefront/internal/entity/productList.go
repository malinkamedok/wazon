package entity

import "github.com/google/uuid"

type ProductList struct {
	UUID uuid.UUID `json:"uuid" `
	Name string    `json:"name"`
}
