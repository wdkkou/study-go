// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type User struct {
	ID    int64
	Name  string
	Email sql.NullString
}
