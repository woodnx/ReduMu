package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/woodnx/ReduMu/api/clock"
	"github.com/woodnx/ReduMu/api/internal/domain/model"
	"github.com/woodnx/ReduMu/api/internal/domain/repository"
	"github.com/woodnx/ReduMu/api/internal/usecase"

	"github.com/google/uuid"
)

func TestAddTask(t *testing.T) {
	t.Parallel()
	clocker := clock.FixedClocker{}
	now := clocker.Now()

	wantTask := &model.Task{
		ID:       model.TaskID(uuid.New()),
		Title:    "test title",
		Status:   model.TaskStatusTodo,
		Deadline: now.Add(24 * time.Hour),
		Created:  now,
		Modified: now,
	}

	type MockParameter struct {
		in  *model.Task
		err error
	}
	cases := map[string]struct {
		title     string
		want      *model.Task
		mockprm   MockParameter
		expectErr bool
	}{
		"C: able to add a task correctly": {
			title: "test title",
			want:  wantTask,
			mockprm: MockParameter{
				in:  wantTask,
				err: nil,
			},
			expectErr: false,
		},
		"E: return a error from repository": {
			title: "test title",
			mockprm: MockParameter{
				in:  wantTask,
				err: errors.New("repository failure"),
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
				CreateFunc: func(pctx context.Context, task *model.Task) error {
					if ctx != pctx {
						t.Fatalf("not want context %v", pctx)
					}

					// usecaseで定義した task と本来入力したい task が異なる
					if d := cmp.Diff(task, tc.mockprm.in); len(d) != 0 {
						t.Fatalf("differs: (-got +want)\n%s", d)
					}
					return tc.mockprm.err
				},
			}

			uc := usecase.NewAddTask(mockRepo)
			got, err := uc.Exec(ctx, tc.title)

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
