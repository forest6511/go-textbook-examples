# 第12章 REST APIの構築

「Go言語の教科書」第12章のサンプルコードです。

## 実行方法

```bash
go run .
```

サーバーが `http://localhost:8080` で起動します。SQLite データベース `books.db` が自動生成されます。

## 動作確認

別のターミナルから以下のコマンドを実行してください。

```bash
# 書籍の作成
curl -s -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Go入門","author":"田中","price":1980}'

# 一覧取得
curl -s http://localhost:8080/books

# ID指定で取得
curl -s http://localhost:8080/books/1

# 削除
curl -s -X DELETE http://localhost:8080/books/1
```

## 停止方法

`Ctrl+C` でサーバーを停止します。

## ファイル構成

```
ch12-rest-api/
├── go.mod           # モジュール定義（modernc.org/sqlite 依存）
├── go.sum           # 依存チェックサム
├── main.go          # サーバー起動・DB初期化・ルーティング
├── model.go         # データモデル（Book, CreateBookRequest）
├── handler.go       # HTTPハンドラ・バリデーション・ヘルパー
├── repository.go    # データベース操作（CRUD）
└── README.md        # このファイル
```

## ポイント

- `modernc.org/sqlite`（pure Go SQLite ドライバ）でデータ永続化
- `database/sql` の共通インターフェースで DB 操作
- 拡張ルーティング（`"GET /books/{id}"`）
- ハンドラをクロージャとして返すパターンで `*sql.DB` を渡す
- SQL プレースホルダ（`?`）で SQL インジェクション防止
- `errors.Is(err, sql.ErrNoRows)` で 404 判定
