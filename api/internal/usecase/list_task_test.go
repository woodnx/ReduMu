package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/woodnx/ReduMu/api/clock"
	"github.com/woodnx/ReduMu/api/internal/domain/model"
	"github.com/woodnx/ReduMu/api/internal/domain/repository"
)

func TestListTask(t *testing.T) {
	t.Parallel()

	clocker := clock.FixedClocker{}
	now := clocker.Now()

	task1 := &model.Task{
		ID:       model.TaskID(uuid.New()),
		Title:    "Task 1",
		Status:   model.TaskStatusTodo,
		Deadline: now.Add(24 * time.Hour),
		Created:  now,
		Modified: now,
	}
	task2 := &model.Task{
		ID:       model.TaskID(uuid.New()),
		Title:    "Task 2",
		Status:   model.TaskStatusDoing,
		Deadline: now.Add(48 * time.Hour),
		Created:  now,
		Modified: now,
	}

	type MockParameter struct {
		out model.Tasks
		err error
	}
	cases := map[string]struct {
		want      model.Tasks
		mockprm   MockParameter
		expectErr bool
	}{
		"C: able to fetch all tasks correctly": {
			want: model.Tasks{task1, task2},
			mockprm: MockParameter{
				out: model.Tasks{task1, task2},
				err: nil,
			},
			expectErr: false,
		},
		"E: return a error from repository": {
			want: nil,
			mockprm: MockParameter{
				out: nil,
				err: errors.New("repository error"),
			},
			expectErr: true,
		},
	}

	for name, tc := range cases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()

			mockRepo := &repository.ITaskRepoMock{
				GetAllFunc: func(pctx context.Context) (model.Tasks, error) {
					if ctx != pctx {
						t.Fatalf("not want context %v", pctx)
					}
					return tc.mockprm.out, tc.mockprm.err
				},
			}

			uc := NewListTask(mockRepo)
			got, err := uc.Exec(ctx)

			if tc.expectErr {
				if err == nil {
					t.Fatal("expected error, but got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if d := cmp.Diff(got, tc.want); len(d) != 0 {
				t.Errorf("task output mismatch: (-got +want)\n%s", d)
			}
		})
	}
}
