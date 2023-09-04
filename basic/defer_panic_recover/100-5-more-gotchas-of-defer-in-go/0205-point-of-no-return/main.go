package main

import (
	"errors"
	"fmt"
)

// #5 — Point of no return?

// 从 deferred func 返回值实际上对调用方没有影响。但是，你仍然可以使用命名返回值来修改返回值。

// Example
func release() error {
	defer func() error {
		return errors.New("error")
	}()
	return nil
}

func main() {
	fmt.Println(release())
}

// Output:
// <nil>

// Solution
// defer 可以修改命名返回值
func release2() (err error) {
	defer func() {
		err = errors.New("error")
	}()
	return nil
}

// Output:
// error

// Why does it work?
// 这里，我们分配一个新值来释放 defer 中的 func's err 返回值。并且，函数返回该值。这里 defer 不会直接返回值，它有助于返回该值。

// Tip
// 你不需要再任何地方都是用 defer。你也可以不使用 defer 直接返回错误。
// 当有多个 return paths 并且你希望在一个位置捕获所有返回路径时，defer 会更方便。最好多考虑下如何进一步简化代码。
