package model

import "gopkg.in/guregu/null.v4"

// Book
type Book struct {
	ID     null.Int
	Title  null.String
	Author null.String
}
