package usecase

import (
	"context"
	"fmt"

	"github.com/woodnx/ReduMu/api/internal/domain/model"
	"github.com/woodnx/ReduMu/api/internal/domain/repository"
)

type AddTask struct {
	repo repository.ITaskRepo
}

func NewAddTask(repo repository.ITaskRepo) *AddTask {
	return &AddTask{repo: repo}
}

func (uc *AddTask) Exec(ctx context.Context, title string) (*model.Task, error) {
	t := &model.Task{
		Title:  title,
		Status: model.TaskStatusTodo,
	}
	err := uc.repo.Create(ctx, t)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return t, nil
}
