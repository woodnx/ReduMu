package main

import (
	"context"

	"github.com/woodnx/ReduMu/api/config"
	api "github.com/woodnx/ReduMu/api/gen"
	"github.com/woodnx/ReduMu/api/internal/handler"
	"github.com/woodnx/ReduMu/api/internal/infra"
)

func NewOgen(ctx context.Context, cfg *config.Config) (*api.Server, func(), error) {
	db, cleanup, err := infra.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	h := handler.New(db)
	srv, err := api.NewServer(h)
	if err != nil {
		return nil, nil, err
	}
	return srv, cleanup, nil
}
