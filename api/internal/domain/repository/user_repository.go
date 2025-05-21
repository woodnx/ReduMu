package repository

import (
	"context"

	"github.com/woodnx/ReduMu/api/internal/domain/model"
)

//go:generate go run github.com/matryer/moq@latest -out ./mock_user_repository.go . IUserRepo
type IUserRepo interface {
	Create(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id model.UserID) (*model.User, error)
	GetAll(ctx context.Context) (model.Users, error)
}
