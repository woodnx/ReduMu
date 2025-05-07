package repository

import (
	"context"

	"github.com/woodnx/ReduMu/api/internal/domain/model"
)

//go:generate go run github.com/matryer/moq@latest -out ./mock_user_repository.go . ITaskRepo
type ITaskRepo interface {
	Create(ctx context.Context, task *model.Task) error
	FindByID(ctx context.Context, id model.TaskID) (*model.Task, error)
	GetAll(ctx context.Context) (model.Tasks, error)
}
