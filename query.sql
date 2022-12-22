-- name: CreateUser :one
INSERT INTO "user" (name, email, username)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM "user" WHERE email = $1;

-- name: GetUserByUsername :one
SELECT * FROM "user" WHERE username = $1;

-- name: GetUserByID :one
SELECT * FROM "user" WHERE id = $1;

-- name: GetUserByEmailOrUsername :one
SELECT * FROM "user" WHERE email = $1 OR username = $2;