package handler

import (
	"database/sql"

	api "github.com/woodnx/ReduMu/api/gen"
	"github.com/woodnx/ReduMu/api/internal/infra"
	"github.com/woodnx/ReduMu/api/internal/usecase"
)

func NewHandler(db *sql.DB) api.Handler {
	tr := infra.NewTaskInfra(db)

	at := usecase.NewAddTask(tr)
	lt := usecase.NewListTask(tr)
	th := NewTaskHandler(at, lt)

	return th
}
