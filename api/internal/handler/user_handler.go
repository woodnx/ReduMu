package handler

import (
	"context"

	api "github.com/woodnx/ReduMu/api/gen"
	"github.com/woodnx/ReduMu/api/internal/usecase"
)

type UserHander struct {
	api.Handler
	registerUser *usecase.RegisterUser
}

func (h *UserHander) UsersPost(ctx context.Context, req api.OptUsersPostReq) error {

	return nil
}
