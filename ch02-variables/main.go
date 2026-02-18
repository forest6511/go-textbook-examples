package main

import "fmt"

// 曜日を定数で定義（iota の例）
const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

// ビットフラグ（iota の応用例）
const (
	ReadPerm  = 1 << iota // 1 (001)
	WritePerm             // 2 (010)
	ExecPerm              // 4 (100)
)

func main() {
	// --- 変数宣言 ---
	var name string = "Go"
	var year int = 2009
	var version = 1.22
	var count int // ゼロ値: 0
	fmt.Println(name, year, version, count)

	// --- 短縮変数宣言 ---
	lang := "Go"
	pi := 3.14
	active := true
	fmt.Println(lang, pi, active)

	// --- ゼロ値の確認 ---
	var i int
	var f float64
	var s string
	var b bool
	fmt.Printf("int: %d, float64: %f, string: %q, bool: %t\n",
		i, f, s, b)

	// --- 型変換 ---
	n := 42
	fv := float64(n)
	uv := uint(fv)
	fmt.Println(n, fv, uv)
	fmt.Printf("型: %T, %T, %T\n", n, fv, uv)

	// --- 定数と iota ---
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Saturday:", Saturday)

	perm := ReadPerm | WritePerm
	fmt.Printf("Read+Write: %d (%03b)\n", perm, perm)

	// --- 演算子 ---
	a, bv := 17, 5
	fmt.Println("加算:", a+bv)
	fmt.Println("除算:", a/bv)
	fmt.Println("剰余:", a%bv)
}
