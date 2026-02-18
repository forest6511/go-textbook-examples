package main

import (
	"database/sql"
)

func allBooks(db *sql.DB) ([]Book, error) {
	rows, err := db.Query(
		"SELECT id, title, author, price FROM books",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		err := rows.Scan(
			&b.ID, &b.Title, &b.Author, &b.Price,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, rows.Err()
}

func bookByID(
	db *sql.DB, id int64,
) (Book, error) {
	var b Book
	err := db.QueryRow(
		"SELECT id, title, author, price"+
			" FROM books WHERE id = ?", id,
	).Scan(&b.ID, &b.Title, &b.Author, &b.Price)
	if err != nil {
		return b, err
	}
	return b, nil
}

func createBook(
	db *sql.DB, req CreateBookRequest,
) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO books (title, author, price)"+
			" VALUES (?, ?, ?)",
		req.Title, req.Author, req.Price,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func deleteBook(db *sql.DB, id int64) error {
	result, err := db.Exec(
		"DELETE FROM books WHERE id = ?", id,
	)
	if err != nil {
		return err
	}
	// 0行影響なら対象不在を呼び出し元に伝える
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return sql.ErrNoRows
	}
	return nil
}
