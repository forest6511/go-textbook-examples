package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/forest6511/go-textbook-examples/ch13-bookmark-app/internal/model"
	"github.com/forest6511/go-textbook-examples/ch13-bookmark-app/internal/repository"
)

// Handler は HTTP リクエストを処理する。
type Handler struct {
	repo *repository.BookmarkRepository
}

// New は Handler を生成する。
func New(
	repo *repository.BookmarkRepository,
) *Handler {
	return &Handler{repo: repo}
}

type errorResponse struct {
	Error string `json:"error"`
}

func writeJSON(
	w http.ResponseWriter,
	status int, data any,
) {
	w.Header().Set(
		"Content-Type", "application/json",
	)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(
	w http.ResponseWriter,
	status int, message string,
) {
	writeJSON(w, status, errorResponse{
		Error: message,
	})
}

// Routes はエンドポイントを mux に登録する。
func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /bookmarks",
		h.listBookmarks)
	mux.HandleFunc("POST /bookmarks",
		h.createBookmark)
	mux.HandleFunc("GET /bookmarks/{id}",
		h.getBookmark)
	mux.HandleFunc("DELETE /bookmarks/{id}",
		h.deleteBookmark)
}

func (h *Handler) createBookmark(
	w http.ResponseWriter, r *http.Request,
) {
	var req model.CreateBookmarkRequest
	if err := json.NewDecoder(r.Body).
		Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest,
			"無効なJSON")
		return
	}
	if req.URL == "" || req.Title == "" {
		writeError(w, http.StatusBadRequest,
			"url と title は必須です")
		return
	}
	bm, err := h.repo.Create(req)
	if err != nil {
		writeError(w,
			http.StatusInternalServerError,
			"登録に失敗しました")
		return
	}
	writeJSON(w, http.StatusCreated, bm)
}

func (h *Handler) listBookmarks(
	w http.ResponseWriter, r *http.Request,
) {
	bookmarks, err := h.repo.All()
	if err != nil {
		writeError(w,
			http.StatusInternalServerError,
			"取得に失敗しました")
		return
	}
	writeJSON(w, http.StatusOK, bookmarks)
}

func (h *Handler) getBookmark(
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
	bm, err := h.repo.FindByID(id)
	if errors.Is(err, sql.ErrNoRows) {
		writeError(w, http.StatusNotFound,
			"ブックマークが見つかりません")
		return
	}
	if err != nil {
		writeError(w,
			http.StatusInternalServerError,
			"取得に失敗しました")
		return
	}
	writeJSON(w, http.StatusOK, bm)
}

func (h *Handler) deleteBookmark(
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
	err = h.repo.Delete(id)
	if errors.Is(err, sql.ErrNoRows) {
		writeError(w, http.StatusNotFound,
			"ブックマークが見つかりません")
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
