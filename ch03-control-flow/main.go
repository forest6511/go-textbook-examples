package main

import "fmt"

func main() {
	// --- if 文 ---
	score := 85
	if score >= 80 {
		fmt.Println("優秀です")
	}

	// --- for 文 (3部分) ---
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// --- range over int (Go 1.22+) ---
	for i := range 5 {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// --- range でスライスを反復 ---
	fruits := []string{"りんご", "みかん", "ぶどう"}
	for i, fruit := range fruits {
		fmt.Printf("%d: %s\n", i, fruit)
	}

	// --- break と continue ---
	for i := range 10 {
		if i == 3 {
			continue
		}
		if i == 7 {
			break
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// --- switch 文 ---
	day := "水曜日"
	switch day {
	case "月曜日":
		fmt.Println("週の始まり")
	case "水曜日":
		fmt.Println("週の半ば")
	default:
		fmt.Println("その他の曜日")
	}

	// --- defer ---
	fmt.Println("開始")
	defer fmt.Println("遅延実行")
	fmt.Println("処理中")
}
