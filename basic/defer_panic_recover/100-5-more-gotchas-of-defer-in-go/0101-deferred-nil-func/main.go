package main

import "fmt"

// #1 — Deferred nil func

// 如果 defer 函数的求值结果为 nil，在包围它的函数结束时会引发 panic，而不是在调用 defer 时

// Example
func test() {
	var run func() = nil
	defer run()
	fmt.Println("runs")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	test()
}

// Output:
// runs
// runtime error: invalid memory address or nil pointer dereference

// Why
// 当函数执行到最后 deferred func 会执行并引发 panic，因为它是 nil。然而 run() 可以被注册而不会出现问题，因为它只在包含的函数结束后才会被调用。
