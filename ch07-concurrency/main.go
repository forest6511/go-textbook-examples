package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// --- goroutine の基本 ---

func sayHello(name string) {
	fmt.Printf("Hello, %s\n", name)
}

// --- sync.WaitGroup ---

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // panic しても確実にカウンタを減算
	fmt.Printf("Worker %d 開始\n", id)
	fmt.Printf("Worker %d 完了\n", id)
}

// --- channel: producer / consumer ---

func producer(out chan<- int) {
	for i := range 5 {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for v := range in {
		fmt.Printf("受信: %d\n", v)
	}
}

// --- sync.Mutex ---

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

// --- context ---

func ctxWorker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: キャンセル\n", id)
			return
		default:
			fmt.Printf("Worker %d: 処理中\n", id)
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func slowTask(ctx context.Context) error {
	select {
	case <-time.After(2 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	// goroutine の基本
	fmt.Println("=== goroutine の基本 ===")
	go sayHello("Alice")
	go sayHello("Bob")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main 終了")

	// sync.WaitGroup
	fmt.Println("\n=== sync.WaitGroup ===")
	var wg sync.WaitGroup
	for i := range 3 {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("全ワーカー完了")

	// channel の基本
	fmt.Println("\n=== channel の基本 ===")
	ch := make(chan string)
	go func() {
		ch <- "Hello from goroutine"
	}()
	msg := <-ch
	fmt.Println(msg)

	// バッファ付き channel
	fmt.Println("\n=== バッファ付き channel ===")
	bufCh := make(chan int, 3)
	bufCh <- 10
	bufCh <- 20
	bufCh <- 30
	fmt.Println(<-bufCh)
	fmt.Println(<-bufCh)
	fmt.Println(<-bufCh)

	// producer / consumer パイプライン
	fmt.Println("\n=== パイプライン ===")
	pipeCh := make(chan int)
	go producer(pipeCh)
	consumer(pipeCh)

	// select
	fmt.Println("\n=== select ===")
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(30 * time.Millisecond)
		ch2 <- "two"
	}()
	for range 2 {
		select {
		case m := <-ch1:
			fmt.Println("ch1:", m)
		case m := <-ch2:
			fmt.Println("ch2:", m)
		}
	}

	// タイムアウト
	fmt.Println("\n=== タイムアウト ===")
	toCh := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		toCh <- "結果"
	}()
	select {
	case m := <-toCh:
		fmt.Println("受信:", m)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("タイムアウト")
	}

	// sync.Mutex
	fmt.Println("\n=== sync.Mutex ===")
	c := SafeCounter{v: make(map[string]int)}
	var wg2 sync.WaitGroup
	for range 1000 {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			c.Inc("key")
		}()
	}
	wg2.Wait()
	fmt.Println("count:", c.Value("key"))

	// context.WithCancel
	fmt.Println("\n=== context.WithCancel ===")
	ctx, cancel := context.WithCancel(
		context.Background(),
	)
	go ctxWorker(ctx, 1)
	go ctxWorker(ctx, 2)
	time.Sleep(120 * time.Millisecond)
	cancel() // 全 goroutine にキャンセルを通知
	time.Sleep(50 * time.Millisecond)
	fmt.Println("cancel 完了")

	// context.WithTimeout
	fmt.Println("\n=== context.WithTimeout ===")
	ctx2, cancel2 := context.WithTimeout(
		context.Background(),
		500*time.Millisecond,
	)
	defer cancel2() // リソースリークを防ぐ
	err := slowTask(ctx2)
	if err != nil {
		fmt.Println("エラー:", err)
	}
}
