# 第8章 テスト

「Go言語の教科書」第8章のサンプルコードです。

## テスト実行

```bash
go test -v ./...
```

## ベンチマーク実行

```bash
go test -bench=. -benchmem
```

## カバレッジ確認

```bash
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## 内容

- テーブル駆動テスト（`TestAdd`, `TestDivide`）
- `t.Run` によるサブテスト
- `t.Parallel` による並列テスト
- `BenchmarkAdd` / `BenchmarkDivide`（テーブル駆動ベンチマーク）
- `ExampleAdd`（Example テスト）
- `t.Helper` によるテストヘルパー

## ファイル構成

```
ch08-testing/
├── go.mod          # モジュール定義
├── calc.go         # テスト対象の関数
├── calc_test.go    # テストファイル
└── README.md       # このファイル
```
