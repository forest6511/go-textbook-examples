package calc

import (
	"fmt"
	"testing"
)

// --- 基本のテスト ---

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 2, 3, 5},
		{"zero", 0, 0, 0},
		{"negative", -1, -2, -3},
		{"mixed", -1, 3, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d",
					tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// --- エラーケースのテスト ---

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{"normal", 10, 3, 3.333, false},
		{"by_one", 10, 1, 10.0, false},
		{"zero_div", 10, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Divide(%g, %g) error = %v, "+
					"wantErr %v", tt.a, tt.b, err, tt.wantErr)
			}
			if !tt.wantErr && (got-tt.want) > 0.01 {
				t.Errorf("Divide(%g, %g) = %g, want %g",
					tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// --- t.Parallel ---

func TestAddParallel(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 2, 3, 5},
		{"negative", -1, -2, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // 並列実行を宣言
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d",
					tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// --- ベンチマーク ---

func BenchmarkAdd(b *testing.B) {
	for b.Loop() {
		Add(100, 200)
	}
}

func BenchmarkDivide(b *testing.B) {
	benchmarks := []struct {
		name string
		a, b float64
	}{
		{"small", 10, 3},
		{"large", 1e18, 7},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for b.Loop() {
				Divide(bm.a, bm.b)
			}
		})
	}
}

// --- Example テスト ---

func ExampleAdd() {
	fmt.Println(Add(2, 3))
	fmt.Println(Add(-1, 1))
	// Output:
	// 5
	// 0
}

// --- t.Helper ---

func assertEqual(t *testing.T, got, want int) {
	t.Helper() // エラー位置を呼び出し元に設定
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAddWithHelper(t *testing.T) {
	assertEqual(t, Add(2, 3), 5)
	assertEqual(t, Add(0, 0), 0)
}
