package usecase

import (
	"context"
	"fmt"

	"github.com/woodnx/ReduMu/api/internal/domain/model"
	"github.com/woodnx/ReduMu/api/internal/domain/repository"
)

type ListTask struct {
	repo repository.ITaskRepo
}

func NewListTask(repo repository.ITaskRepo) *ListTask {
	return &ListTask{repo: repo}
}

func (uc *ListTask) Exec(ctx context.Context) (model.Tasks, error) {
	ts, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}

	return ts, nil
}
