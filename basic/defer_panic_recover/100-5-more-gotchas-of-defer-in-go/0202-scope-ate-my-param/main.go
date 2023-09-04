package main

import (
	"errors"
	"fmt"
	"io"
)

// #2—Scope ate my param

// 这实际上是一个范围陷阱，但是，我想向你展示它如何使用defer和命名返回值

// Example
// 定义一个函数，它注册一个延迟函数，以便在函数返回后延迟释放资源 r。
func release(r io.Closer) (err error) {
	defer func() {
		if err := r.Close(); err != nil { // creates a new err variable
			// ..
		}
	}()
	// ..
	return
}

// 创建一个在 close 时返回 error 的 reader
type reader struct{}

func (r reader) Close() error {
	return errors.New("Close Error")
}

// Let's try
// reader 调用其 Close() 方法时始终返回 error，release() 将在 defer 中调用它。

func main() {
	r := reader{}
	err := release(r)
	fmt.Print(err)
}

// Output:
// <nil>

// err 是 nil，但是我们期待的是：“Close Error”
// Why is that?
// 延迟函数 if 块内的赋值用新的 err 变量修改 命名返回值 err。因此，release() 返回原始 err 返回值。

// Solution
// 在 release() 中使用原始 err 变量。此代码不声明新 err 变量，而是使用原始 err 返回值(使用 = 代替 :=)。这解决了 shadowing 问题。
func release2(r io.Closer) (err error) {
	defer func() {
		if err = r.Close(); err != nil { // overwrites err
			// ..
		}
	}()
	// ..
	return
}
