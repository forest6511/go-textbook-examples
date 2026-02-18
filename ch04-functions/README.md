# 第4章 関数とメソッド

「Go言語の教科書」第4章のサンプルコードです。

## 実行方法

```bash
go run main.go
```

## 内容

- 複数の戻り値と `(T, error)` パターン
- 名前付き戻り値（naked return）
- 可変長引数（`...`）
- 高階関数（`apply`、`sort.Slice`）
- 関数型に名前を付ける（`type Transformer func(int) int`）
- クロージャ（状態のキャプチャ）
- メソッド（構造体・基本型）
- 値レシーバとポインタレシーバ
- `String()` メソッド（`fmt.Stringer`）

## ファイル構成

```
ch04-functions/
├── go.mod      # モジュール定義
├── main.go     # メインプログラム
└── README.md   # このファイル
```
