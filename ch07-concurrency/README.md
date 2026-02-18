# 第7章 並行処理

「Go言語の教科書」第7章のサンプルコードです。

## 実行方法

```bash
go run main.go
```

レースディテクタを有効にして実行する場合:

```bash
go run -race main.go
```

## 注意事項

並行処理の出力は実行ごとに順序が異なる場合があります。これは goroutine のスケジューリングが非決定的であるためで、正常な動作です。

## 内容

- goroutine の起動（`go` キーワード）
- `sync.WaitGroup` による完了待ち
- channel の基本（送受信、バッファ付き）
- channel の方向指定（`chan<-`、`<-chan`）
- producer / consumer パイプラインパターン
- `select` による複数 channel の待ち受け
- `time.After` によるタイムアウト
- `sync.Mutex` による排他制御（`SafeCounter`）
- `context.WithCancel` によるキャンセル
- `context.WithTimeout` によるタイムアウト

## ファイル構成

```
ch07-concurrency/
├── go.mod      # モジュール定義
├── main.go     # メインプログラム
└── README.md   # このファイル
```
