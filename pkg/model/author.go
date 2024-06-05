package model

import "time"

type Author struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Bio        string    `json:"bio"`
	Books      []*Book   `json:"books"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	AddedBy    *User     `json:"added_by"`
}
