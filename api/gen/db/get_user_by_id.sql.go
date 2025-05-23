// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: get_user_by_id.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const getUserByID = `-- name: GetUserByID :one
SELECT 
    "id",
    "name", 
    "password", 
    "role", 
    "created", 
    "modified"
FROM "user"
WHERE "id" == $1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Password,
		&i.Role,
		&i.Created,
		&i.Modified,
	)
	return i, err
}
