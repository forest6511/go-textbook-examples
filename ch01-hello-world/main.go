package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World!")

	now := time.Now()
	fmt.Println("現在時刻:", now.Format("2006-01-02 15:04:05"))
}
