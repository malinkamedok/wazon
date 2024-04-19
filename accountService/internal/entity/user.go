package entity

import "github.com/google/uuid"

type User struct {
	Id      uuid.UUID `json:"id,omitempty"`
	Name    string    `json:"name"`
	SurName string    `json:"surname"`
	Login   string    `json:"login"`
}
