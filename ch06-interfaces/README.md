# 第6章 インターフェースとエラー処理

「Go言語の教科書」第6章のサンプルコードです。

## 実行方法

```bash
go run main.go
```

## 内容

- インターフェースの定義と暗黙的な実装（`Greeter`）
- `fmt.Stringer` インターフェース（`Color`）
- 型アサーションと型スイッチ
- `error` インターフェースとエラー処理
- `fmt.Errorf` + `%w` によるエラーのラップ
- `errors.Is` / `errors.As` によるエラーの検査
- カスタムエラー型（`ValidationError`, `HTTPError` + `Unwrap`）
- `defer` の3つのルール
- `panic` と `recover`

## ファイル構成

```
ch06-interfaces/
├── go.mod      # モジュール定義
├── main.go     # メインプログラム
└── README.md   # このファイル
```
