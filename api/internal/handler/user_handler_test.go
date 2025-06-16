package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	api "github.com/woodnx/ReduMu/api/gen"
	"github.com/woodnx/ReduMu/api/testutil"
)

func TestUsersPost(t *testing.T) {
	type want struct {
		status  int
		rspFile string
	}
	cases := map[string]struct {
		reqFile string
		want    want
	}{
		"ok": {
			reqFile: "testdata/register_user/ok_req.json",
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/register_user/ok_rsp.json",
			},
		},
		// "badRequest": {
		// 	reqFile: "testdata/register_user/bas_req.json",
		// 	want: want{
		// 		status:  http.StatusBadRequest,
		// 		rspFile: "testdata/register_user/bad_rsp.json",
		// 	},
		// },
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			sqlDB := testutil.SetupTestDB(t)
			h := New(sqlDB)
			srv, err := api.NewServer(h)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/users", bytes.NewReader(testutil.LoadFile(t, tc.reqFile)))
			srv.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tc.want.status, testutil.LoadFile(t, tc.want.rspFile),
			)
		})
	}
}
