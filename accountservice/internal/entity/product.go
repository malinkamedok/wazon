package entity

import "github.com/google/uuid"

type Product struct {
	Id          uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
}

type User struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	SurName string `json:"surname"`
	Login   string `json:"login"`
}
