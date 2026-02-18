package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "modernc.org/sqlite"

	"github.com/forest6511/go-textbook-examples/ch13-bookmark-app/internal/model"
	"github.com/forest6511/go-textbook-examples/ch13-bookmark-app/internal/repository"
)

func setupTestHandler(
	t *testing.T,
) (*Handler, *http.ServeMux) {
	t.Helper()
	// ファイルではなくメモリ上にDBを作成
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Close() })
	repo := repository.New(db)
	if err := repo.InitTable(); err != nil {
		t.Fatal(err)
	}
	h := New(repo)
	mux := http.NewServeMux()
	h.Routes(mux)
	return h, mux
}

func TestCreateBookmark(t *testing.T) {
	_, mux := setupTestHandler(t)

	body := strings.NewReader(
		`{"url":"https://go.dev",` +
			`"title":"Go公式サイト"}`,
	)
	req := httptest.NewRequest(
		"POST", "/bookmarks", body,
	)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d",
			rec.Code, http.StatusCreated)
	}
	var bm model.Bookmark
	json.NewDecoder(rec.Body).Decode(&bm)
	if bm.URL != "https://go.dev" {
		t.Errorf("url = %q, want %q",
			bm.URL, "https://go.dev")
	}
}

func TestCreateBookmark_validation(
	t *testing.T,
) {
	_, mux := setupTestHandler(t)

	tests := []struct {
		name   string
		body   string
		status int
	}{
		{"URL空",
			`{"url":"","title":"T"}`, 400},
		{"Title空",
			`{"url":"https://x","title":""}`,
			400},
		{"不正JSON", `{bad`, 400},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(
				"POST", "/bookmarks",
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
}

func TestBookmarkFlow(t *testing.T) {
	_, mux := setupTestHandler(t)

	// 登録
	body := strings.NewReader(
		`{"url":"https://go.dev","title":"Go"}`,
	)
	req := httptest.NewRequest(
		"POST", "/bookmarks", body,
	)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("create: status = %d",
			rec.Code)
	}

	// 一覧取得で1件あることを確認
	req = httptest.NewRequest(
		"GET", "/bookmarks", nil,
	)
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	var list []model.Bookmark
	json.NewDecoder(rec.Body).Decode(&list)
	if len(list) != 1 {
		t.Fatalf("len = %d, want 1",
			len(list))
	}
}
