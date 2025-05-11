-- name: GetAllUsers :many
SELECT 
    "id",
    "name", 
    "password", 
    "role", 
    "created", 
    "modified"
FROM "user";