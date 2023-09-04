package main

import (
	"fmt"
	"net/http"
)

// #2— Calling defer in the wrong order
// #2— 以错误的顺序调用 defer

// Example
// 当 http.Get 失败时，此代码将 panic
func do() error {
	res, err := http.Get("http://notexists")
	defer res.Body.Close()
	if err != nil {
		return err
	}

	// ..code...

	return nil
}

func main() {
	do()
	x, y := 1, 2
	defer func(a int) {
		fmt.Println("defer x, y = ", a, y)
	}(x)

	x += 100
	y += 100
	fmt.Println("x, y = ", x, y)
	//do2
}

// output:
// panic: runtime error: invalid memory address or nil pointer dereference

// Why?
// 因为这里我们没有检查请求是否成功。这里它失败了，我们在 nil变量(res) 上调用 Body，因此 panic

// Solution
// 始终在成功分配资源后使用 defer。对于此示例这意味着：仅在 http.Get 成功时使用 defer
func do2() error {
	res, err := http.Get("https://notexists")
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return err
	}

	// ..code...

	return nil
}

// 在上述的代码中，当出现错误时，代码将返回错误，否则，当函数在 defer 中返回时 它将关闭 res.Body。

// Side-Note
// 在这里，你还需要检查 resp 是否为 nil，这是 http.Get 中的警告。通常，当出现错误时，响应将为 nil 并返回 error。
// 但是，当你收到 redirection error 时，响应不会返回 nil，但会返回 error。上面的代码，确保关闭了 response body，如果你不打算使用它，你还需要丢弃收到的数据。

// Update:
// 这个问题在 http 中看起来并不存在，我需要为此找到一个更好的例子。所以，它可能对某些代码仍然有效，但不是 http。
