package gateway

import (
	"database/sql"

	_ "github.com/lib/pq"
	sqlc "github.com/Hermes-chat-App/hermes-auth-server/internal/db"
)

var Queries *sqlc.Queries

func Init(connStr string) error {
	var err error

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	Queries = sqlc.New(db)

	return nil
}
