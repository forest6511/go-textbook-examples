package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

// --- JSON ---

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	// ファイル操作
	fmt.Println("=== ファイル操作 ===")
	content := []byte("Hello, Go!\n")
	err := os.WriteFile("output.txt", content, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "書き込みエラー: %v\n", err)
		os.Exit(1)
	}
	data, err := os.ReadFile("output.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "読み込みエラー: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(string(data))
	os.Remove("output.txt") // テスト後にクリーンアップ

	// JSON エンコード
	fmt.Println("=== JSON エンコード ===")
	user := User{
		Name: "田中", Email: "tanaka@example.com", Age: 30,
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	fmt.Println(string(jsonData))

	// JSON デコード
	fmt.Println("\n=== JSON デコード ===")
	jsonStr := `{"name":"佐藤","email":"sato@example.com","age":25}`
	var u User
	err = json.Unmarshal([]byte(jsonStr), &u)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	fmt.Printf("%s（%d歳）\n", u.Name, u.Age)

	// 文字列操作
	fmt.Println("\n=== 文字列操作 ===")
	s := "Hello, Go World"
	fmt.Println(strings.Contains(s, "Go"))
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.Replace(s, "World", "言語", 1))
	fmt.Println(strings.Split("a,b,c", ","))
	fmt.Println(strings.TrimSpace("  hello  "))
	fmt.Println(strings.HasPrefix(s, "Hello"))

	// strconv
	fmt.Println("\n=== strconv ===")
	n, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	fmt.Println(n + 8)
	fmt.Println("値: " + strconv.Itoa(100))
	f, err := strconv.ParseFloat("3.14", 64)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	fmt.Printf("%.2f\n", f)

	// 時刻操作
	fmt.Println("\n=== 時刻操作 ===")
	now := time.Now()
	fmt.Println("現在時刻:", now.Format("2006-01-02 15:04:05"))
	later := now.Add(1*time.Hour + 30*time.Minute)
	fmt.Println("1時間30分後:", later.Format("15:04:05"))

	t, err := time.Parse("2006-01-02", "2026-03-15")
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	fmt.Println("解析:", t.Format("2006/01/02"))

	// 構造化ログ
	fmt.Println("\n=== log/slog ===")
	slog.Info("サーバー起動",
		"host", "localhost",
		"port", 8080,
	)
	slog.Warn("ディスク使用率が高い",
		"usage", 85.5,
		"threshold", 80.0,
	)
}
