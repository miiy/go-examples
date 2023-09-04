package main

import "fmt"

// #4 — Defer in a block

// 你可能期望 deferred func 在块结束后运行，但它没有，它仅在包含的func结束后执行。这也适用于所有块：for, switch 等，除了我们在前面的陷阱中看到的函数块。
// 因为：defer属于一个函数而不是一个块

// Example
func main() {
	{
		defer func() {
			fmt.Println("block: defer runs")
		}()
		fmt.Println("block: ends")
	}
	fmt.Println("main: ends")
}

// Output:
// block: ends
// main: ends
// block: defer runs

// Explanation
// 上面的 deferred func 仅在函数结束时运行，而不是在 deferred func 包围的块结束时运行，如代码所示，你可以仅使用大括号创建单独的块。

// Another Solution
// 如果你想在一个块中运行defer，你可以把它转换成一个func，比如匿名函数，
func main2() {
	func() {
		defer func() {
			fmt.Println("func: defer runs")
		}()
		fmt.Println("func: ends")
	}()
	fmt.Println("main: ends")
}
