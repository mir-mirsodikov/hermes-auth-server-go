// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Username string
	Verfied  sql.NullBool
}

type Verification struct {
	UserID    uuid.UUID
	Code      int32
	CreatedAt sql.NullTime
}
