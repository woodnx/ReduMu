package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/woodnx/ReduMu/api/clock"
	api "github.com/woodnx/ReduMu/api/gen"
	"github.com/woodnx/ReduMu/api/testutil"
)

func TestTasksGet(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	now := clock.FixedClocker{}.Now()

	type want struct {
		status  int
		rspFile string
	}
	cases := map[string]struct {
		tasks []api.Task
		want  want
	}{
		"C: Able to fetch two tasks": {
			tasks: []api.Task{
				{
					ID:        uuid.MustParse("93fa7c3e-322e-5cb5-1158-5696ff2ff9ce"),
					Title:     "Task 1",
					Deadline:  now.Add(24 * time.Hour),
					CreatedAt: now,
					UpdatedAt: api.OptDate{Value: now, Set: true},
				},
				{
					ID:        uuid.MustParse("e09c18b9-fe14-5128-0b38-c264acd00544"),
					Title:     "Task 2",
					Deadline:  now.Add(48 * time.Hour),
					CreatedAt: now,
					UpdatedAt: api.OptDate{Value: now, Set: true},
				},
			},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_task/ok_rsp.json",
			},
		},
		"C: Able to fetch zero tasks": {
			tasks: []api.Task{},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_task/empty_rsp.json",
			},
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			sqlDB := testutil.SetupTestDB(t)

			// INSERT
			for _, task := range tc.tasks {
				_, err := sqlDB.ExecContext(ctx,
					`INSERT INTO task (id, title, deadline, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`,
					task.ID, task.Title, task.Deadline, task.CreatedAt, task.UpdatedAt.Value,
				)
				require.NoError(t, err)
			}

			// handler & サーバ構築
			h := New(sqlDB)
			srv, err := api.NewServer(h)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/tasks", nil)
			srv.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tc.want.status, testutil.LoadFile(t, tc.want.rspFile),
			)
		})
	}
}
