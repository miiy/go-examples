package main

import (
	"os"
)

// #3— Not checking for errors
// #3— 不检查错误

// 仅将清理逻辑委托给defer，并不意味着将毫无问题地释放。你还会错过可能有用的错误信息，并失去通过沉没它们来诊断隐藏问题的能力。

// Not Good
// 这里，f.Close() 可能会返回错误，但我们不会意识到它
func do() error {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	// ...code...

	return nil
}

func main() {
	do()
}

// Better
// 最好检查错误，而不是委托和忘记。
// 你可以通过将 defer 内的代码放到一个辅助函数中来简化下面的代码，这样做可能有点凌乱，但可以帮助你理解问题。
func do2() error {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			// log etc
		}
	}()

	// ...code...

	return nil
}

// Better
// 你还可以使用命名返回值在 defer 中返回 error
func do3() (err error) {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}

	defer func() {
		if ferr := f.Close(); ferr != nil {
			err = ferr
		}
	}()

	// ...code...

	return nil
}

// Side-Note
// 你还可以使用 github.com/pkg/errors 包装多个错误，这可能是必要的，因为，f.Close 在 defer 中可能会吞下它之前的任何错误。
// 通过将错误与另一个错误包装，会将此信息写入 log，以便你可以使用更多的数据诊断错误。
