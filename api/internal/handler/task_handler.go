package handler

import (
	"context"

	"github.com/google/uuid"
	api "github.com/woodnx/ReduMu/api/gen"
	"github.com/woodnx/ReduMu/api/internal/usecase"
)

type TaskHander struct {
	api.Handler
	addTask  *usecase.AddTask
	listTask *usecase.ListTask
}

func NewTaskHandler(at *usecase.AddTask, lt *usecase.ListTask) *TaskHander {
	return &TaskHander{
		addTask:  at,
		listTask: lt,
	}
}

func (h *TaskHander) TasksGet(ctx context.Context) ([]api.Task, error) {
	ts, err := h.listTask.Exec(ctx)
	if err != nil {
		return nil, err
	}

	tasks := []api.Task{}
	for _, t := range ts {
		task := api.Task{
			ID:        uuid.UUID(t.ID),
			Title:     t.Title,
			Deadline:  t.Deadline,
			CreatedAt: t.Created,
			UpdatedAt: api.OptDate{Value: t.Modified, Set: true},
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
