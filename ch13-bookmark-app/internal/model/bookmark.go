package model

import "time"

// Bookmark はブックマークの永続化モデル。
type Bookmark struct {
	ID        int64     `json:"id"`
	URL       string    `json:"url"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateBookmarkRequest は登録リクエストの形式。
type CreateBookmarkRequest struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}
