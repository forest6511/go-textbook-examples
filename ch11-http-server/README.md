# 第11章 HTTPサーバーの基本

「Go言語の教科書」第11章のサンプルコードです。

## 実行方法

```bash
go run .
```

サーバーが `http://localhost:8080` で起動します。

## 動作確認

別のターミナルから以下のコマンドを実行してください。

```bash
# Hello エンドポイント
curl http://localhost:8080/hello

# 検索（クエリパラメータ）
curl "http://localhost:8080/search?q=Go+HTTP&page=2"

# ユーザー取得（JSON レスポンス）
curl http://localhost:8080/users/1

# ユーザー作成（JSON リクエスト）
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"田中","email":"tanaka@example.com"}'
```

## 停止方法

`Ctrl+C` でグレースフルシャットダウンします（処理中のリクエスト完了後に停止）。

## ファイル構成

```
ch11-http-server/
├── go.mod           # モジュール定義
├── main.go          # サーバー起動・ルーティング・シャットダウン
├── handler.go       # ハンドラ関数
├── middleware.go     # ミドルウェア（ログ・リカバリ）
└── README.md        # このファイル
```

## ポイント

- Go 1.22 の拡張ルーティング（`"GET /path/{param}"`）を使用
- `http.NewServeMux()` で独自ルーターを作成（DefaultServeMux を避ける）
- ミドルウェアチェーン: recover → logging → ハンドラ
- `http.Server` 構造体でタイムアウトを設定
- シグナル受信によるグレースフルシャットダウン
