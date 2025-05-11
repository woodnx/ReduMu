-- name: AddUser :one
INSERT INTO "user" (
   "id", "name", "password", "role", "created", "modified"
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;