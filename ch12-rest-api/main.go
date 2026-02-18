package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	_ "modernc.org/sqlite"
)

func openDB(path string) *sql.DB {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		slog.Error("DB接続失敗", "error", err)
		os.Exit(1)
	}
	// sql.Open は遅延接続のため、実際に通信できるか検証する
	if err := db.Ping(); err != nil {
		slog.Error("DB Ping失敗", "error", err)
		os.Exit(1)
	}
	return db
}

func initTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS books (
		id     INTEGER PRIMARY KEY AUTOINCREMENT,
		title  TEXT NOT NULL,
		author TEXT NOT NULL,
		price  INTEGER NOT NULL
	)`
	if _, err := db.Exec(query); err != nil {
		slog.Error("テーブル作成失敗", "error", err)
		os.Exit(1)
	}
}

func main() {
	db := openDB("books.db")
	defer db.Close()
	initTable(db)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /books",
		handleListBooks(db))
	mux.HandleFunc("POST /books",
		handleCreateBook(db))
	mux.HandleFunc("GET /books/{id}",
		handleGetBook(db))
	mux.HandleFunc("DELETE /books/{id}",
		handleDeleteBook(db))

	slog.Info("サーバー起動", "addr", ":8080")
	if err := http.ListenAndServe(
		":8080", mux); err != nil {
		slog.Error("サーバーエラー", "error", err)
		os.Exit(1)
	}
}
