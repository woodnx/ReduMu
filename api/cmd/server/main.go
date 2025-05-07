package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/woodnx/ReduMu/api/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// 環境変数の取得
	cfg, err := config.New()
	if err != nil {
		return err
	}

	// ネットワークリスナーの作成
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	srv, cleanup, err := NewOgen(ctx, cfg)
	if err != nil {
		return err
	}
	defer cleanup()
	s := NewServer(l, srv)
	return s.Run(ctx)
}
