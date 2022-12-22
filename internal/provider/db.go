package provider

import (
	"database/sql"

	sqlc "github.com/Hermes-chat-App/hermes-auth-server/internal/db"
	_ "github.com/lib/pq"
)

var Queries *sqlc.Queries

func DBInit(connStr string) error {
	var err error

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	Queries = sqlc.New(db)

	return nil
}
