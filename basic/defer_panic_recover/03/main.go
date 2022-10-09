package main

import "time"

func main() {
	defer println("in main")
	go func() {
		defer println("in goroutine")
		panic("")
	}()

	time.Sleep(1 * time.Second)
}

//in goroutine
//panic:

// defer 只会出发当前 Goroutine 的延迟函数调用
