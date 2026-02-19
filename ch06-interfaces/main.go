package main

import (
	"errors"
	"fmt"
)

// --- インターフェース ---

type Greeter interface {
	Greet() string
}

type Japanese struct{ Name string }

func (j Japanese) Greet() string {
	return "こんにちは、" + j.Name + "です"
}

type English struct{ Name string }

func (e English) Greet() string {
	return "Hello, I'm " + e.Name
}

func printGreeting(g Greeter) {
	fmt.Println(g.Greet())
}

// --- fmt.Stringer ---

type Color struct {
	R, G, B uint8
}

func (c Color) String() string {
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

// --- 型アサーションと型スイッチ ---

func describe(i any) {
	s, ok := i.(string)
	if ok {
		fmt.Printf("文字列: %q（長さ %d）\n", s, len(s))
	} else {
		fmt.Printf("文字列ではない: %v\n", i)
	}
}

func classify(i any) string {
	switch v := i.(type) {
	case int:
		return fmt.Sprintf("整数: %d", v)
	case string:
		return fmt.Sprintf("文字列: %q", v)
	case bool:
		return fmt.Sprintf("真偽値: %t", v)
	default:
		return fmt.Sprintf("不明: %v", v)
	}
}

// --- エラー処理 ---

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("ゼロ除算はできません")
	}
	return a / b, nil
}

// --- エラーのラップ ---

var ErrNotFound = errors.New("not found")

func findUser(id int) error {
	return fmt.Errorf("user %d: %w", id, ErrNotFound)
}

// --- カスタムエラー型 ---

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func validate(name string) error {
	if name == "" {
		return &ValidationError{
			Field: "name", Message: "必須項目です",
		}
	}
	return nil
}

// --- HTTPError with Unwrap ---

type HTTPError struct {
	Code    int
	Message string
	Err     error
}

func (e *HTTPError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("HTTP %d: %s: %s",
			e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("HTTP %d: %s", e.Code, e.Message)
}

func (e *HTTPError) Unwrap() error {
	return e.Err
}

// --- defer ---

func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("パニック捕捉: %v", r)
		}
	}()
	return a / b, nil
}

func main() {
	// インターフェース
	fmt.Println("=== インターフェース ===")
	printGreeting(Japanese{Name: "田中"})
	printGreeting(English{Name: "Alice"})

	// fmt.Stringer
	fmt.Println("\n=== fmt.Stringer ===")
	red := Color{R: 255, G: 0, B: 0}
	fmt.Println(red)

	// 型アサーション
	fmt.Println("\n=== 型アサーション ===")
	describe("hello")
	describe(42)

	// 型スイッチ
	fmt.Println("\n=== 型スイッチ ===")
	fmt.Println(classify(42))
	fmt.Println(classify("Go"))
	fmt.Println(classify(true))

	// エラー処理
	fmt.Println("\n=== エラー処理 ===")
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

	// エラーのラップ
	fmt.Println("\n=== エラーのラップ ===")
	err = findUser(42)
	fmt.Println(err)
	fmt.Println("Is ErrNotFound:", errors.Is(err, ErrNotFound))

	// errors.AsType（Go 1.26+: 型安全なエラー検査）
	fmt.Println("\n=== errors.AsType ===")
	err = validate("")
	if ve, ok := errors.AsType[*ValidationError](err); ok {
		fmt.Printf("バリデーションエラー: %s\n", ve.Field)
	}

	// HTTPError with Unwrap
	fmt.Println("\n=== カスタムエラー型 ===")
	cause := errors.New("connection refused")
	httpErr := &HTTPError{
		Code: 503, Message: "サービス利用不可",
		Err: cause,
	}
	fmt.Println(httpErr)
	fmt.Println("Is cause:", errors.Is(httpErr, cause))

	// defer のルール
	fmt.Println("\n=== defer ルール1 ===")
	x := 10
	defer fmt.Println("defer:", x)
	x = 20
	fmt.Println("main:", x)

	// panic と recover
	fmt.Println("\n=== panic と recover ===")
	r, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Println("結果:", r)
	}
}
