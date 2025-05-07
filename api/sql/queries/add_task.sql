-- name: AddTask :one
INSERT INTO "task" (
    "id",
    "title", 
    "status", 
    "created", 
    "modified"
) 
VALUES ($1, $2, $3, $4, $5)
RETURNING *;