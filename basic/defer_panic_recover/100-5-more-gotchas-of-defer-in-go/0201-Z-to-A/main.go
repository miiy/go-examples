package main

import "fmt"

// #1 — Z to A

func main() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

// Output:
// 3210

// Go runtime 将 deferred funcs 保存在栈中，这意味着他们以相反的顺序工作。
