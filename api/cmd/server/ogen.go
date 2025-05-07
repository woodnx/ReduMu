package main

import (
	"context"

	"github.com/woodnx/ReduMu/api/config"
	api "github.com/woodnx/ReduMu/api/gen"
	"github.com/woodnx/ReduMu/api/internal/handler"
	"github.com/woodnx/ReduMu/api/internal/infra"
	"github.com/woodnx/ReduMu/api/internal/usecase"
)

func NewOgen(ctx context.Context, cfg *config.Config) (*api.Server, func(), error) {
	db, cleanup, err := infra.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	tr := infra.NewTaskInfra(db)
	at := usecase.NewAddTask(tr)
	lt := usecase.NewListTask(tr)
	th := handler.NewTaskHandler(at, lt)

	srv, err := api.NewServer(th)
	if err != nil {
		return nil, nil, err
	}
	return srv, cleanup, nil
}
