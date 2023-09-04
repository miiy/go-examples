package main

import "fmt"

// 失效的崩溃恢复

func main() {
	defer println("in main")
	if err := recover(); err != nil {
		fmt.Println(err)
	}

	panic("unknown err")
}

// Output:
// in main
// panic: unknown err

// recover 只有在发生 panic 之后调用才会生效
// 所以我们需要在 defer 中使用 recover 关键字

// Solution
func main2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("unknown err")
}
