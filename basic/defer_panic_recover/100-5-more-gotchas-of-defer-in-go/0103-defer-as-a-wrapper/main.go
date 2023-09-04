package main

import "fmt"

// #3 — Defer as a wrapper
// 有时你需要结合闭包使用 defer，以更实用，或者我现在无法猜测的其他一些原因。例如，打开数据库连接，然后运行一些查询，并确保在最后断开连接。

// Example
type database struct{}

func (db *database) connect() (disconnect func()) {
	fmt.Println("connect")
	return func() {
		fmt.Println("disconnect")
	}
}

func main() {
	db := &database{}
	defer db.connect()

	fmt.Println("query db...")
}

// Output:
// query db...
// connect

// Why it doesn't work?
// 它不会断开连接，而是在最后连接，这是一个错误。这里发生的唯一一件事情就是 connect() 保存到一边，直到 func 结束并且它不会运行。

// Solution
// 现在，db.connect() 返回一个函数，当函数结束时调用 defer 断开数据库连接。
func main2() {
	db := &database{}
	close := db.connect()
	defer close()

	fmt.Println("query db...")
}

// Output:
// connect
// query db...
// disconnect

// Bad Practice
// 虽然这样做是一种不好的实践，但我想向你展示如何在没有变量的情况下做到。所以，我希望你能看到 defer，Go 语言通常的工作原理。
func main3() {
	db := database{}
	defer db.connect()()
	fmt.Println("query db...")
}

// Output:
// connect
// query db...
// disconnect

// 此代码在技术上与上述解决方案几乎相同。在这里，第一个括号用于连接到数据库（立即发生 defer db.connect()），然后第二个括号用于在函数结束时运行disconnector func(the returned closure)
// db.connect()创建一个值，该值是闭包，然后defer注册它。db.connect()的值首先需要解析才能想defer注册，这与defer没有直接关系，但它可能会解决你可能遇到的一些问题。
