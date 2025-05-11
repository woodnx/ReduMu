package repository

import (
	"context"

	"github.com/woodnx/ReduMu/api/internal/domain/model"
)

type IUserRepo interface {
	Create(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id model.UserID) (*model.User, error)
	GetAll(ctx context.Context) (model.Users, error)
}
