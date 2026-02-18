package main

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "modernc.org/sqlite"

	"github.com/forest6511/go-textbook-examples/ch13-bookmark-app/internal/handler"
	"github.com/forest6511/go-textbook-examples/ch13-bookmark-app/internal/repository"
)

func loggingMiddleware(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter,
			r *http.Request,
		) {
			slog.Info("リクエスト受信",
				"method", r.Method,
				"path", r.URL.Path,
			)
			next.ServeHTTP(w, r)
		},
	)
}

func main() {
	db, err := sql.Open("sqlite", "bookmarks.db")
	if err != nil {
		slog.Error("DB接続失敗", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	repo := repository.New(db)
	if err := repo.InitTable(); err != nil {
		slog.Error("テーブル作成失敗",
			"error", err)
		os.Exit(1)
	}

	h := handler.New(repo)
	mux := http.NewServeMux()
	h.Routes(mux)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: loggingMiddleware(mux),
	}

	// Ctrl+C で graceful shutdown を実行
	go func() {
		ctx, stop := signal.NotifyContext(
			context.Background(),
			os.Interrupt,
		)
		defer stop()
		<-ctx.Done()
		slog.Info("シャットダウン開始")
		shutCtx, cancel := context.WithTimeout(
			context.Background(),
			5*time.Second,
		)
		defer cancel()
		srv.Shutdown(shutCtx)
	}()

	slog.Info("サーバー起動", "addr", ":8080")
	if err := srv.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {
		slog.Error("サーバーエラー",
			"error", err)
		os.Exit(1)
	}
}
