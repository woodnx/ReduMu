package usecase

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/woodnx/ReduMu/api/internal/domain/model"
	"github.com/woodnx/ReduMu/api/internal/domain/repository"
)

func TestListTask(t *testing.T) {
	t.Parallel()

	wantID := model.TaskID(uuid.New())
	wantTitle := "test title"
	wantTask := &model.Task{
		ID:    wantID,
		Title: wantTitle,
	}
	mockRepo := &repository.ITaskRepoMock{
		GetAllFunc: func(ctx context.Context) (model.Tasks, error) {
			return model.Tasks{wantTask}, nil
		},
	}

	uc := NewListTask(mockRepo)

	tasks, err := uc.Exec(context.Background())
	assert.NoError(t, err)
	assert.Len(t, tasks, 1)
	assert.Equal(t, "test title", tasks[0].Title)
}
