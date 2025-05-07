-- name: GetAllTasks :many
SELECT 
    "id", 
    "title", 
    "status", 
    "created", 
    "modified"
FROM task;