package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/woodnx/ReduMu/api/internal/domain/model"
	"github.com/woodnx/ReduMu/api/internal/domain/repository"
)

func TestAddTask(t *testing.T) {
	t.Parallel()

	var createdTask *model.Task

	mockRepo := &repository.ITaskRepoMock{
		CreateFunc: func(ctx context.Context, task *model.Task) error {
			createdTask = task
			return nil
		},
	}

	uc := NewAddTask(mockRepo)
	title := "New Task"

	task, err := uc.Exec(context.Background(), title)

	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, title, task.Title)
	assert.Equal(t, model.TaskStatusTodo, task.Status)
	assert.Equal(t, createdTask, task)
}
