-- name: CreateUser :one
INSERT INTO "user" (name, email, username)
VALUES ($1, $2, $3)
RETURNING *;
