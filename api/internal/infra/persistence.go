package infra

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/woodnx/ReduMu/api/clock"
	"github.com/woodnx/ReduMu/api/config"
	"github.com/woodnx/ReduMu/api/gen/db"
)

func New(ctx context.Context, cfg *config.Config) (*sql.DB, func(), error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf(
			"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
			cfg.DBUser, cfg.DBPassword,
			cfg.DBHost, cfg.DBPort,
			cfg.DBName,
		),
	)
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, nil, err
	}
	return db, func() { _ = db.Close() }, nil
}

// // type Beginner interface {
// // 	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
// // }

// // type Preparer interface {
// // 	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
// // }

// // type Execer interface {
// // 	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
// // 	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
// // }

// // type Queryer interface {
// // 	Preparer
// // 	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
// // 	QueryRowxContext(ctx context.Context, query string, args ...any) *sqlx.Row
// // 	GetContext(ctx context.Context, dest interface{}, query string, args ...any) error
// // 	SelectContext(ctx context.Context, dest interface{}, query string, args ...any) error
// // }

// // var (
// // 	_ Beginner = (*sqlx.DB)(nil)
// // 	_ Preparer = (*sqlx.DB)(nil)
// // 	_ Queryer  = (*sqlx.DB)(nil)
// // 	_ Execer   = (*sqlx.DB)(nil)
// // 	_ Execer   = (*sqlx.Tx)(nil)
// // )

type Repository struct {
	Clocker clock.Clocker
	q       *db.Queries
}
