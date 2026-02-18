package repository

import (
	"database/sql"
	"time"

	"github.com/forest6511/go-textbook-examples/ch13-bookmark-app/internal/model"
)

// BookmarkRepository はブックマークの永続化を担当する。
type BookmarkRepository struct {
	db *sql.DB
}

// New は BookmarkRepository を生成する。
func New(db *sql.DB) *BookmarkRepository {
	return &BookmarkRepository{db: db}
}

// InitTable はブックマーク用テーブルを作成する。
func (r *BookmarkRepository) InitTable() error {
	query := `CREATE TABLE IF NOT EXISTS bookmarks (
		id         INTEGER PRIMARY KEY AUTOINCREMENT,
		url        TEXT NOT NULL,
		title      TEXT NOT NULL,
		created_at TEXT NOT NULL
	)`
	_, err := r.db.Exec(query)
	return err
}

// Create はブックマークを登録する。
func (r *BookmarkRepository) Create(
	req model.CreateBookmarkRequest,
) (model.Bookmark, error) {
	now := time.Now().UTC()
	result, err := r.db.Exec(
		`INSERT INTO bookmarks
		 (url, title, created_at)
		 VALUES (?, ?, ?)`,
		req.URL, req.Title,
		now.Format(time.RFC3339),
	)
	if err != nil {
		return model.Bookmark{}, err
	}
	// SQLite は LastInsertId を常にサポートする
	id, _ := result.LastInsertId()
	return model.Bookmark{
		ID: id, URL: req.URL,
		Title: req.Title, CreatedAt: now,
	}, nil
}

// All は全ブックマークを取得する。
func (r *BookmarkRepository) All() (
	[]model.Bookmark, error,
) {
	rows, err := r.db.Query(
		`SELECT id, url, title, created_at
		 FROM bookmarks ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookmarks []model.Bookmark
	for rows.Next() {
		var b model.Bookmark
		var createdAt string
		if err := rows.Scan(
			&b.ID, &b.URL,
			&b.Title, &createdAt,
		); err != nil {
			return nil, err
		}
		// Create で RFC3339 形式に統一しているため
		// パースエラーは発生しない
		b.CreatedAt, _ = time.Parse(
			time.RFC3339, createdAt,
		)
		bookmarks = append(bookmarks, b)
	}
	return bookmarks, rows.Err()
}

// FindByID は指定IDのブックマークを取得する。
func (r *BookmarkRepository) FindByID(
	id int64,
) (model.Bookmark, error) {
	var b model.Bookmark
	var createdAt string
	err := r.db.QueryRow(
		`SELECT id, url, title, created_at
		 FROM bookmarks WHERE id = ?`, id,
	).Scan(
		&b.ID, &b.URL,
		&b.Title, &createdAt,
	)
	if err != nil {
		return model.Bookmark{}, err
	}
	// Create と同じ RFC3339 形式のためパースは安全
	b.CreatedAt, _ = time.Parse(
		time.RFC3339, createdAt,
	)
	return b, nil
}

// Delete は指定IDのブックマークを削除する。
func (r *BookmarkRepository) Delete(
	id int64,
) error {
	result, err := r.db.Exec(
		`DELETE FROM bookmarks WHERE id = ?`,
		id,
	)
	if err != nil {
		return err
	}
	// 0行影響なら対象不在を呼び出し元に伝える
	n, _ := result.RowsAffected()
	if n == 0 {
		return sql.ErrNoRows
	}
	return nil
}
