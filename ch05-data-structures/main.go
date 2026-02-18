package main

import (
	"fmt"
	"sort"
)

// --- スライス ---

func sliceDemo() {
	fmt.Println("=== 配列 ===")
	b := [3]string{"Go", "Rust", "Python"}
	fmt.Println(b)

	fmt.Println("\n=== スライスの作成 ===")
	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Println(fruits)
	nums := make([]int, 3, 5)
	fmt.Println(nums, "長さ:", len(nums), "容量:", cap(nums))

	fmt.Println("\n=== append ===")
	var s []int
	s = append(s, 1)
	s = append(s, 2, 3)
	fmt.Println(s)
	more := []int{4, 5, 6}
	s = append(s, more...)
	fmt.Println(s)

	fmt.Println("\n=== スライス式 ===")
	n := []int{10, 20, 30, 40, 50}
	fmt.Println(n[1:4])
	fmt.Println(n[:3])
	fmt.Println(n[2:])

	fmt.Println("\n=== copy ===")
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copied := copy(dst, src)
	fmt.Println(dst, "コピー数:", copied)
	dst[0] = 999
	fmt.Println("src:", src)
	fmt.Println("dst:", dst)

	fmt.Println("\n=== スライスの内部構造 ===")
	si := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(si), cap(si), si)
	si = append(si, 10)
	fmt.Printf("len=%d cap=%d %v\n", len(si), cap(si), si)
	si = append(si, 20)
	fmt.Printf("len=%d cap=%d %v\n", len(si), cap(si), si)
	si = append(si, 30)
	fmt.Printf("len=%d cap=%d %v\n", len(si), cap(si), si)
}

// --- マップ ---

func mapDemo() {
	fmt.Println("\n=== マップの基本操作 ===")
	ages := map[string]int{"Alice": 30, "Bob": 25}
	ages["Charlie"] = 35
	ages["Alice"] = 31
	fmt.Println(ages["Alice"])
	delete(ages, "Bob")
	fmt.Println(ages)

	fmt.Println("\n=== comma-ok イディオム ===")
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	code, ok := colors["red"]
	fmt.Printf("red: %s, exists: %t\n", code, ok)
	code, ok = colors["yellow"]
	fmt.Printf("yellow: %q, exists: %t\n", code, ok)

	fmt.Println("\n=== マップの反復（ソート済み） ===")
	scores := map[string]int{
		"Charlie": 85, "Alice": 92, "Bob": 78,
	}
	keys := make([]string, 0, len(scores))
	for k := range scores {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, scores[k])
	}
}

// --- 構造体 ---

type User struct {
	Name  string
	Email string
	Age   int
}

type Address struct {
	City    string
	ZipCode string
}

type Employee struct {
	Name    string
	Address // 埋め込み
}

// --- Stack ---

type Stack struct {
	items []int
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	last := len(s.items) - 1
	v := s.items[last]
	s.items = s.items[:last]
	return v, true
}

func (s *Stack) Len() int {
	return len(s.items)
}

// --- ポインタ ---

func increment(n *int) {
	*n++
}

func main() {
	sliceDemo()
	mapDemo()

	fmt.Println("\n=== 構造体 ===")
	u1 := User{Name: "Alice", Email: "alice@example.com", Age: 30}
	fmt.Println(u1)

	fmt.Println("\n=== 埋め込み ===")
	e := Employee{
		Name:    "Tanaka",
		Address: Address{City: "Tokyo", ZipCode: "100-0001"},
	}
	fmt.Println(e.City)

	fmt.Println("\n=== Stack ===")
	var st Stack
	st.Push(10)
	st.Push(20)
	st.Push(30)
	fmt.Println("長さ:", st.Len())
	if v, ok := st.Pop(); ok {
		fmt.Println("Pop:", v)
	}
	fmt.Println("長さ:", st.Len())

	fmt.Println("\n=== ポインタ ===")
	x := 10
	increment(&x)
	fmt.Println(x)

	p := new(int)
	fmt.Println(*p)
	*p = 42
	fmt.Println(*p)
}
