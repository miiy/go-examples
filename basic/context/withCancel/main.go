package main

import (
	"context"
	"fmt"
)
//This example demonstrates the use of a cancelable context to prevent a goroutine leak.
//By the end of the example function, the goroutine started by gen will return without leaking.
//此示例演示了使用可取消的 context 来防止一个 goroutine 泄露。
//到示例函数结束时，由 gen 启动的 goroutine 将返回而不会泄露。
func main() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	// gen 在单独的 goroutine 生成整数并将它们发送到返回的 channel。
	// gen 的调用者需要在消耗完生成的整数后取消 context，以免泄露由 gen 启动的内部 goroutine
	gen := func(ctx context.Context) <- chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <- ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

//Output:
//1
//2
//3
//4
//5
