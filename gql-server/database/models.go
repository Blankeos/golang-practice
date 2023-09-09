// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package database

import (
	"database/sql"
)

type Post struct {
	ID    int64          `json:"id"`
	Title sql.NullString `json:"title"`
	Body  sql.NullString `json:"body"`
}

type Todo struct {
	ID   int64          `json:"id"`
	Text sql.NullString `json:"text"`
	Done sql.NullBool   `json:"done"`
}

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
