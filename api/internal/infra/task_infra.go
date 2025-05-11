package infra

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/woodnx/ReduMu/api/clock"
	"github.com/woodnx/ReduMu/api/gen/db"
	"github.com/woodnx/ReduMu/api/internal/domain/model"
	"github.com/woodnx/ReduMu/api/internal/domain/repository"
)

type taskRepository struct {
	Clocker clock.Clocker
	q       *db.Queries
}

func NewTaskInfra(sqldb *sql.DB) repository.ITaskRepo {
	return &taskRepository{
		Clocker: clock.RealClocker{},
		q:       db.New(sqldb),
	}
}

func (r *taskRepository) GetAll(ctx context.Context) (model.Tasks, error) {
	ts, err := r.q.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	tasks := model.Tasks{}
	for _, t := range ts {
		tasks = append(tasks, &model.Task{
			ID:       model.TaskID(t.ID),
			Status:   model.TaskStatus(t.Status),
			Created:  t.Created,
			Modified: t.Modified,
		})
	}
	return tasks, nil
}

func (r *taskRepository) Create(ctx context.Context, t *model.Task) error {
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	id := uuid.New()

	arg := db.AddTaskParams{
		ID:       id,
		Title:    t.Title,
		Status:   string(t.Status),
		Created:  t.Created,
		Modified: t.Modified,
	}

	_, err := r.q.AddTask(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}

func (r *taskRepository) FindByID(ctx context.Context, id model.TaskID) (*model.Task, error) {
	return nil, nil
}
