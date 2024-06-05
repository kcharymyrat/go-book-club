package model

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	UserName  string
	Email     string
	Password  string
}
