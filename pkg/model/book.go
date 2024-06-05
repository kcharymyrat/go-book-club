package model

import "time"

type Book struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Summary    string    `json:"summary"`
	Year       int       `json:"year"`
	Authors    []*Author `json:"authors"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	AddedBy    *User     `json:"added_by"`
}
