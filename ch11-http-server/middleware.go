package main

import (
	"log/slog"
	"net/http"
)

func loggingMiddleware(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			slog.Info("リクエスト受信",
				"method", r.Method,
				"path", r.URL.Path,
			)
			next.ServeHTTP(w, r)
		},
	)
}

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					slog.Error("パニック発生",
						"error", err,
					)
					http.Error(w,
						"Internal Server Error",
						http.StatusInternalServerError,
					)
				}
			}()
			next.ServeHTTP(w, r)
		},
	)
}
