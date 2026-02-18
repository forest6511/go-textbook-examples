package main

// Book はデータベースのレコードを表す。
type Book struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

// CreateBookRequest はクライアントからの作成リクエスト。
type CreateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}
