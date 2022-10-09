package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("main defer")
	}()
	deferCall()
}

func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
	fmt.Println("test")
}

//打印后
//打印中
//打印前
//main defer
//panic: 触发异常
//
//goroutine 1 [running]:
//main.deferCall()
//		/go-examples/basic/defer_panic_recover/main.go:14 +0x6b
//main.main()
//		/go-examples/basic/defer_panic_recover/main.go:6 +0x17
//
//Process finished with the exit code 2

// defer的顺序是后进先出
// 调用 panic 后会立刻停止当前函数的剩余代码，并在当前 Goroutine 中递归执行调用的 defer
