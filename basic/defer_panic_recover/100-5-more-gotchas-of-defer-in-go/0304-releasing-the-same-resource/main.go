package main

import (
	"fmt"
	"io"
	"os"
)

// #4— Releasing the same resource
// #4— 释放相同的资源

// 如果你尝试使用相同的变量关闭另外一个资源，它可能不会按预期运行。

// Example
// 这个看起来没问题的代码试图关闭同一资源两次。这里，第二个 r 变量将关闭两次。因为，下面的第二个资源将改变 r 变量
func do() error {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("defer close book.txt err %v\n", err)
		}
	}()

	// ...code...

	f, err = os.Open("another-book.txt")
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("defer close another-book.txt err %v\n", err)
		}
	}()

	return nil
}

func main() {
	err := do()
	fmt.Println("error:", err)
	//err := do2()
	//fmt.Println("error:", err)
}

// output:
// defer close book.txt err close another-book.txt: file already closed

// Why
// 当 defer 执行时，只使用最后一个变量。因此，f 变量 会成为最后一个 (another-book.txt)。而且，两个 defer 都会将其视为最后一个。

// Solution
func do2() error {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}

	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			fmt.Printf("defer close book.txt err %v\n", err)
		}
	}(f)

	// ...code...

	f, err = os.Open("another-book.txt")
	if err != nil {
		return err
	}

	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			fmt.Printf("defer close another-book.txt err %v\n", err)
		}
	}(f)

	return nil
}

// Output:
// error: <nil>
