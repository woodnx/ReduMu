package handler

import (
	"context"

	api "github.com/woodnx/ReduMu/api/gen"
)

func (h *Handler) UsersPost(ctx context.Context, req api.OptUsersPostReq) error {
	h.mux.Lock()
	defer h.mux.Unlock()

	h.registerUser.Exec(ctx, req.Value.Name, req.Value.Password, req.Value.Password)
	return nil
}
