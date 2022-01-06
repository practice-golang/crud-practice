package model

import "gopkg.in/guregu/null.v4"

// Book
type Book struct {
	ID     null.Int    `json:"id"     db:"IDX"`
	Title  null.String `json:"title"  db:"TITLE"`
	Author null.String `json:"author" db:"AUTHOR"`
}
