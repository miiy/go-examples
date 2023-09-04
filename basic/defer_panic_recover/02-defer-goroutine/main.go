package main

import "time"

// 跨协程失效
// panic 只会触发当前 Goroutine 的延迟函数调用。

func main() {
	defer println("in main")
	go func() {
		defer println("in goroutine")
		panic("")
	}()
	time.Sleep(1 * time.Second)
}

// Output:
// in goroutine
// panic:

func main2() {
	defer println("in main")
	func() {
		defer println("in goroutine")
		panic("")
	}()
	time.Sleep(1 * time.Second)
}

// Output:
// in goroutine
// in main
// panic:
