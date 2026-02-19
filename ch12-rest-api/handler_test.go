package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "modernc.org/sqlite"
)

// setupTestDB はテスト用の in-memory SQLite を準備する。
func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Close() })
	initTable(db)
	return db
}

// setupMux はテスト用の ServeMux を構築する。
func setupMux(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /books",
		handleListBooks(db))
	mux.HandleFunc("POST /books",
		handleCreateBook(db))
	mux.HandleFunc("GET /books/{id}",
		handleGetBook(db))
	mux.HandleFunc("DELETE /books/{id}",
		handleDeleteBook(db))
	return mux
}

func TestValidateCreateBook(t *testing.T) {
	tests := []struct {
		name    string
		req     CreateBookRequest
		wantErr bool
	}{
		{"valid",
			CreateBookRequest{
				Title: "Go入門", Author: "田中",
				Price: 1980,
			}, false},
		{"empty_title",
			CreateBookRequest{
				Title: "", Author: "田中",
				Price: 1980,
			}, true},
		{"empty_author",
			CreateBookRequest{
				Title: "Go入門", Author: "",
				Price: 1980,
			}, true},
		{"negative_price",
			CreateBookRequest{
				Title: "Go入門", Author: "田中",
				Price: -1,
			}, true},
		{"zero_price",
			CreateBookRequest{
				Title: "Go入門", Author: "田中",
				Price: 0,
			}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCreateBook(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"validateCreateBook(%+v) error = %v, "+
						"wantErr %v",
					tt.req, err, tt.wantErr)
			}
		})
	}
}

func TestHandleListBooks(t *testing.T) {
	db := setupTestDB(t)
	mux := setupMux(db)

	// 空の状態で一覧取得
	req := httptest.NewRequest("GET", "/books", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d",
			rec.Code, http.StatusOK)
	}
	var books []Book
	json.NewDecoder(rec.Body).Decode(&books)
	if len(books) != 0 {
		t.Errorf("len = %d, want 0", len(books))
	}
}

func TestHandleCreateBook(t *testing.T) {
	db := setupTestDB(t)
	mux := setupMux(db)

	tests := []struct {
		name   string
		body   string
		status int
	}{
		{"valid",
			`{"title":"Go入門","author":"田中","price":1980}`,
			http.StatusCreated},
		{"invalid_json",
			`{bad`,
			http.StatusBadRequest},
		{"empty_title",
			`{"title":"","author":"田中","price":1980}`,
			http.StatusBadRequest},
		{"negative_price",
			`{"title":"Go入門","author":"田中","price":-1}`,
			http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(
				"POST", "/books",
				strings.NewReader(tt.body),
			)
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)
			if rec.Code != tt.status {
				t.Errorf("status = %d, want %d",
					rec.Code, tt.status)
			}
		})
	}

	// 正常作成後のレスポンス検証
	body := strings.NewReader(
		`{"title":"実践Go","author":"鈴木","price":2480}`,
	)
	req := httptest.NewRequest("POST", "/books", body)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	var book Book
	json.NewDecoder(rec.Body).Decode(&book)
	if book.Title != "実践Go" {
		t.Errorf("title = %q, want %q",
			book.Title, "実践Go")
	}
	if book.Price != 2480 {
		t.Errorf("price = %d, want %d",
			book.Price, 2480)
	}
}

func TestHandleGetBook(t *testing.T) {
	db := setupTestDB(t)
	mux := setupMux(db)

	// まず1冊登録
	body := strings.NewReader(
		`{"title":"Go入門","author":"田中","price":1980}`,
	)
	req := httptest.NewRequest("POST", "/books", body)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	var created Book
	json.NewDecoder(rec.Body).Decode(&created)

	tests := []struct {
		name   string
		path   string
		status int
	}{
		{"existing", "/books/1",
			http.StatusOK},
		{"not_found", "/books/999",
			http.StatusNotFound},
		{"invalid_id", "/books/abc",
			http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(
				"GET", tt.path, nil,
			)
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)
			if rec.Code != tt.status {
				t.Errorf("status = %d, want %d",
					rec.Code, tt.status)
			}
		})
	}

	// 正常取得時のレスポンス検証
	req = httptest.NewRequest("GET", "/books/1", nil)
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	var got Book
	json.NewDecoder(rec.Body).Decode(&got)
	if got.Title != "Go入門" {
		t.Errorf("title = %q, want %q",
			got.Title, "Go入門")
	}
}

func TestHandleDeleteBook(t *testing.T) {
	db := setupTestDB(t)
	mux := setupMux(db)

	// 1冊登録
	body := strings.NewReader(
		`{"title":"Go入門","author":"田中","price":1980}`,
	)
	req := httptest.NewRequest("POST", "/books", body)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	tests := []struct {
		name   string
		path   string
		status int
	}{
		{"existing", "/books/1",
			http.StatusNoContent},
		{"already_deleted", "/books/1",
			http.StatusNotFound},
		{"not_found", "/books/999",
			http.StatusNotFound},
		{"invalid_id", "/books/abc",
			http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(
				"DELETE", tt.path, nil,
			)
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)
			if rec.Code != tt.status {
				t.Errorf("status = %d, want %d",
					rec.Code, tt.status)
			}
		})
	}
}
