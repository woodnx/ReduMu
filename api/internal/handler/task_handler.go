package handler

import (
	"context"

	"github.com/google/uuid"
	api "github.com/woodnx/ReduMu/api/gen"
)

func (h *Handler) TasksGet(ctx context.Context) ([]api.Task, error) {
	h.mux.Lock()
	defer h.mux.Unlock()

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
