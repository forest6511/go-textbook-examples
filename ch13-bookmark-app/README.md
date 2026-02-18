# ch13-bookmark-app

第13章「実践: ブックマーク管理アプリ」のサンプルコード。

第11章（HTTPサーバー）と第12章（REST API）で学んだ技術を組み合わせ、
`cmd/` + `internal/` のプロダクション構成でブックマーク管理APIを構築します。

## ディレクトリ構成

```
ch13-bookmark-app/
├── cmd/server/main.go          # エントリーポイント
├── internal/
│   ├── handler/handler.go      # HTTPハンドラ
│   ├── handler/handler_test.go # ハンドラテスト
│   ├── model/bookmark.go       # データモデル
│   └── repository/bookmark.go  # DB操作
├── go.mod
└── go.sum
```

## 実行方法

```bash
go run ./cmd/server/
```

サーバーが `:8080` で起動します。Ctrl+C で graceful shutdown します。

## エンドポイント

| メソッド | パス | 説明 |
|---------|------|------|
| POST | /bookmarks | ブックマーク登録 |
| GET | /bookmarks | 一覧取得 |
| GET | /bookmarks/{id} | 個別取得 |
| DELETE | /bookmarks/{id} | 削除 |

## 使用例

```bash
# 登録
curl -X POST http://localhost:8080/bookmarks \
  -d '{"url":"https://go.dev","title":"Go公式サイト"}'

# 一覧取得
curl http://localhost:8080/bookmarks

# 個別取得
curl http://localhost:8080/bookmarks/1

# 削除
curl -X DELETE http://localhost:8080/bookmarks/1
```

## テスト

```bash
go test ./...
```

## 依存パッケージ

- [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite) — Pure Go SQLite ドライバ（cgo 不要）

## 対応する書籍の章

「Go言語の教科書」第13章
