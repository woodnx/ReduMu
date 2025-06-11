package testutil

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	pgContainer *postgres.PostgresContainer
	dbConn      *sql.DB
	setupOnce   sync.Once
	setupErr    error
)

// SetupTestDB はテスト用のDBコンテナを起動し、マイグレーション済みの*sql.DBと
// クリーンアップ用の関数を返します。
// この実装では、テスト実行全体でコンテナを1度だけ起動し、再利用します。
func SetupTestDB(t *testing.T) *sql.DB {
	t.Helper()

	// sync.Once を使って、並列テストでもコンテナのセットアップが1度しか実行されないようにする
	setupOnce.Do(func() {
		ctx := context.Background()

		pgContainer, setupErr = postgres.Run(ctx,
			"postgres:15-alpine",
			postgres.WithDatabase("test-db"),
			postgres.WithUsername("user"),
			postgres.WithPassword("password"),
			testcontainers.WithWaitStrategy(
				wait.ForLog("database system is ready to accept connections").
					WithOccurrence(2).
					WithStartupTimeout(10*time.Second),
			),
		)
		if setupErr != nil {
			return
		}

		// コンテナへの接続文字列を取得
		conStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
		if err != nil {
			setupErr = fmt.Errorf("failed to get connection string: %w", err)
			return
		}

		// マイグレーションの実行
		migrationsPath := filepath.Join("..", "..", "sql", "migrations") // パスの起点をプロジェクトルートに合わせる
		m, err := migrate.New(fmt.Sprintf("file://%s", migrationsPath), conStr)
		if err != nil {
			setupErr = fmt.Errorf("failed to create migrate instance: %w", err)
			return
		}
		if err = m.Up(); err != nil {
			setupErr = fmt.Errorf("failed to run migrate up: %w", err)
			return
		}

		// DB接続を確立
		db, err := sql.Open("postgres", conStr)
		if err != nil {
			setupErr = fmt.Errorf("failed to open database connection: %w", err)
			return
		}

		dbConn = db
	})

	// setupOnce.Do の中でエラーが発生していたらテストを失敗させる
	require.NoError(t, setupErr)

	// 各テストの前にテーブルをクリーンアップする
	// これにより、各テストは独立したクリーンな状態で開始できる
	t.Cleanup(func() {
		// ここでは例として tasks テーブルをTRUNCATEする
		_, err := dbConn.Exec("TRUNCATE TABLE tasks RESTART IDENTITY")
		require.NoError(t, err)
	})

	return dbConn
}
