-- name: GetUserByID :one
SELECT 
    "id",
    "name", 
    "password", 
    "role", 
    "created", 
    "modified"
FROM "user"
WHERE "id" == $1;