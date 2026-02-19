package main

import (
	"fmt"
	"math"
	"slices"
)

// --- 複数の戻り値 ---

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("ゼロ除算はできません")
	}
	return a / b, nil
}

// --- 名前付き戻り値 ---

func swap(a, b int) (x, y int) {
	x = b
	y = a
	return // x, y の現在値がそのまま返される（naked return）
}

// --- 可変長引数 ---

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// --- 高階関数 ---

func apply(nums []int, f func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = f(n)
	}
	return result
}

// --- 関数型に名前を付ける ---

type Transformer func(int) int

func applyAll(n int, fns ...Transformer) []int {
	results := make([]int, len(fns))
	for i, f := range fns {
		results[i] = f(n)
	}
	return results
}

// --- クロージャ ---

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// --- メソッド（構造体） ---

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// --- メソッド（構造体以外の型） ---

type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	return float64(c)*9/5 + 32
}

func (c Celsius) String() string {
	return fmt.Sprintf("%.1f°C", float64(c))
}

// --- 値レシーバとポインタレシーバ ---

type Counter struct {
	Value int
}

func (c Counter) Current() int {
	return c.Value
}

func (c *Counter) Increment() {
	c.Value++
}

func (c *Counter) Reset() {
	c.Value = 0
}

func main() {
	fmt.Println("=== 複数の戻り値 ===")
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", result)
	}

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("エラー:", err)
	}

	fmt.Println("\n=== 名前付き戻り値 ===")
	a, b := swap(1, 2)
	fmt.Println(a, b)

	fmt.Println("\n=== 可変長引数 ===")
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(10, 20, 30, 40, 50))
	numbers := []int{100, 200, 300}
	fmt.Println(sum(numbers...))

	fmt.Println("\n=== 高階関数 ===")
	nums := []int{1, 2, 3, 4, 5}
	doubled := apply(nums, func(n int) int { return n * 2 })
	fmt.Println("doubled:", doubled)
	squared := apply(nums, func(n int) int { return n * n })
	fmt.Println("squared:", squared)

	fmt.Println("\n=== sort.Slice ===")
	names := []string{"Charlie", "Alice", "Bob"}
	slices.Sort(names)
	fmt.Println(names)

	fmt.Println("\n=== 関数型に名前を付ける ===")
	double := Transformer(func(n int) int { return n * 2 })
	square := Transformer(func(n int) int { return n * n })
	fmt.Println(applyAll(5, double, square))

	fmt.Println("\n=== クロージャ ===")
	c := counter()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	c2 := counter()
	fmt.Println(c2())

	fmt.Println("\n=== メソッド（Circle） ===")
	circle := Circle{Radius: 5.0}
	fmt.Printf("面積: %.2f\n", circle.Area())
	fmt.Printf("周長: %.2f\n", circle.Perimeter())

	fmt.Println("\n=== メソッド（Celsius） ===")
	temp := Celsius(100)
	fmt.Println(temp)
	fmt.Printf("華氏: %.1f°F\n", temp.ToFahrenheit())

	fmt.Println("\n=== 値レシーバとポインタレシーバ ===")
	ct := Counter{Value: 0}
	ct.Increment()
	ct.Increment()
	ct.Increment()
	fmt.Println("現在値:", ct.Current())
	ct.Reset()
	fmt.Println("リセット後:", ct.Current())
}
