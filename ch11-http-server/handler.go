package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User はユーザー情報を表す。
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUserRequest はユーザー作成リクエストのボディ。
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ErrorResponse はエラー時の統一レスポンス形式。
type ErrorResponse struct {
	Error string `json:"error"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go!")
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	keyword := query.Get("q")
	page := query.Get("page")
	if page == "" {
		page = "1"
	}
	fmt.Fprintf(w, "検索: %s（ページ %s）\n", keyword, page)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID: 1, Name: "田中", Email: "tanaka@example.com",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "無効なJSON")
		return
	}

	// 空の名前はDB保存時にエラーになるため事前チェック
	if req.Name == "" {
		writeError(w, http.StatusBadRequest, "name は必須です")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "ユーザー作成: " + req.Name,
	})
}

func writeError(
	w http.ResponseWriter,
	status int,
	message string,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error: message,
	})
}
