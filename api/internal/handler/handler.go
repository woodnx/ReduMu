package handler

import (
	"database/sql"
	"sync"

	api "github.com/woodnx/ReduMu/api/gen"
	"github.com/woodnx/ReduMu/api/internal/infra"
	"github.com/woodnx/ReduMu/api/internal/usecase"
)

type Handler struct {
	api.UnimplementedHandler
	mux          sync.Mutex
	addTask      *usecase.AddTask
	listTask     *usecase.ListTask
	registerUser *usecase.RegisterUser
}

func New(db *sql.DB) *Handler {
	tr := infra.NewTaskInfra(db)
	ur := infra.NewUserInfra(db)

	addTask := usecase.NewAddTask(tr)
	listTask := usecase.NewListTask(tr)
	registerUser := usecase.NewRegisterUser(ur)

	return &Handler{
		addTask:      addTask,
		listTask:     listTask,
		registerUser: registerUser,
	}
}
