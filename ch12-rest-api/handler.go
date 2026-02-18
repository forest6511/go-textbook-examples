package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// ErrorResponse はエラー時の統一レスポンス形式。
type ErrorResponse struct {
	Error string `json:"error"`
}

func writeJSON(
	w http.ResponseWriter, status int, data any,
) {
	w.Header().Set(
		"Content-Type", "application/json",
	)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(
	w http.ResponseWriter,
	status int,
	message string,
) {
	writeJSON(w, status, ErrorResponse{
		Error: message,
	})
}

func validateCreateBook(
	req CreateBookRequest,
) error {
	if req.Title == "" || req.Author == "" {
		return fmt.Errorf(
			"title と author は必須です")
	}
	if req.Price < 0 {
		return fmt.Errorf(
			"price は0以上にしてください")
	}
	return nil
}

func handleListBooks(
	db *sql.DB,
) http.HandlerFunc {
	return func(
		w http.ResponseWriter, r *http.Request,
	) {
		books, err := allBooks(db)
		if err != nil {
			writeError(w,
				http.StatusInternalServerError,
				"データ取得に失敗しました")
			return
		}
		writeJSON(w, http.StatusOK, books)
	}
}

func handleCreateBook(
	db *sql.DB,
) http.HandlerFunc {
	return func(
		w http.ResponseWriter, r *http.Request,
	) {
		var req CreateBookRequest
		if err := json.NewDecoder(r.Body).
			Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest,
				"無効なJSON")
			return
		}
		if err := validateCreateBook(req); err != nil {
			writeError(w, http.StatusBadRequest,
				err.Error())
			return
		}
		id, err := createBook(db, req)
		if err != nil {
			writeError(w,
				http.StatusInternalServerError,
				"作成に失敗しました")
			return
		}
		book := Book{
			ID: id, Title: req.Title,
			Author: req.Author, Price: req.Price,
		}
		writeJSON(w, http.StatusCreated, book)
	}
}

func handleGetBook(
	db *sql.DB,
) http.HandlerFunc {
	return func(
		w http.ResponseWriter, r *http.Request,
	) {
		id, err := strconv.ParseInt(
			r.PathValue("id"), 10, 64,
		)
		if err != nil {
			writeError(w, http.StatusBadRequest,
				"無効なID")
			return
		}
		book, err := bookByID(db, id)
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound,
				"書籍が見つかりません")
			return
		}
		if err != nil {
			writeError(w,
				http.StatusInternalServerError,
				"取得に失敗しました")
			return
		}
		writeJSON(w, http.StatusOK, book)
	}
}

func handleDeleteBook(
	db *sql.DB,
) http.HandlerFunc {
	return func(
		w http.ResponseWriter, r *http.Request,
	) {
		id, err := strconv.ParseInt(
			r.PathValue("id"), 10, 64,
		)
		if err != nil {
			writeError(w, http.StatusBadRequest,
				"無効なID")
			return
		}
		err = deleteBook(db, id)
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound,
				"書籍が見つかりません")
			return
		}
		if err != nil {
			writeError(w,
				http.StatusInternalServerError,
				"削除に失敗しました")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
