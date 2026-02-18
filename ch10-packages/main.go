package main

import (
	"fmt"

	"github.com/forest6511/go-textbook-examples/ch10-packages/internal/greeting"
)

func main() {
	// internal パッケージのエクスポート関数を呼び出す
	msg := greeting.Hello("Go")
	fmt.Println(msg)
}
