# 第10章 パッケージ設計とモジュール

「Go言語の教科書」第10章のサンプルコードです。

## 実行方法

```bash
go run main.go
```

## 内容

- `internal` パッケージによるアクセス制御の実例
- マルチパッケージ構成のプロジェクト

## ファイル構成

```
ch10-packages/
├── go.mod
├── main.go                    # メインプログラム
├── internal/
│   └── greeting/
│       └── greeting.go        # internal パッケージ
└── README.md                  # このファイル
```

## ポイント

- `internal/greeting` パッケージはこのモジュール内からのみインポートできます
- `greeting.Hello()` はエクスポート関数（大文字で始まる）
- `greeting.formatName()` は非エクスポート関数（小文字で始まる）で、外部からは呼べません
