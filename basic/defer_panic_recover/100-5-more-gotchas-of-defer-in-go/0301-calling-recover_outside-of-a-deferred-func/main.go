package main

import "fmt"

// #1 — Calling recover outside of a deferred func
// #1 - 在 deferred func 之外调用 recover

// 你应该始终在 deferred func 中调用 recover()。
// 当 panic 时，在 defer 之外调用 recovery() 不会捕获它并且 recover() 将返回 nil

// Example
func do() {
	recover()
	panic("error")
}

func main() {
	do()
}

// Output:
// panic: error

// Solution
// 只需要在 defer 中使用 recover()，你就可以防止此问题。
func do2() {
	defer func() {
		r := recover()
		fmt.Println("recovered:", r)
	}()
	panic("error")
}

// Output:
// recovered: error
