package testutil

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
)

func OpenDBForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	port := 5432
	if _, defined := os.LookupEnv("CI"); defined {
		port = 5432
	}
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("postgres:postgres@tcp(db:%d)/redumu?parseTime=true", port),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(
		func() { _ = db.Close() },
	)
	return sqlx.NewDb(db, "postgres")
}
