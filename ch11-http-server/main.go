package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", helloHandler)
	mux.HandleFunc("GET /search", searchHandler)
	mux.HandleFunc("GET /users/{id}", getUser)
	mux.HandleFunc("POST /users", createUser)

	// 内側から: mux → logging → recover
	handler := recoverMiddleware(
		loggingMiddleware(mux),
	)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// メインスレッドをブロックせずシグナルを監視するため
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigCh
		slog.Info("シャットダウン開始", "signal", sig)

		ctx, cancel := context.WithTimeout(
			context.Background(), 10*time.Second,
		)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("シャットダウン失敗", "error", err)
		}
	}()

	slog.Info("サーバー起動", "addr", srv.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		slog.Error("サーバーエラー", "error", err)
		os.Exit(1)
	}
	slog.Info("サーバー停止完了")
}
