package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
	}()
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(1 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}

// output:
// 2
// 1

// runtime.NumGoroutine() 返回当前存在的 goroutine 数
