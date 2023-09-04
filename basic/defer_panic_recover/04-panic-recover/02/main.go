package main

// 嵌套崩溃
// Go 语言中的 panic 是可以多次嵌套调用的。
func main() {
	defer println("in main")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()
	panic("panic once")
}

// Output:
// in main
// panic: panic once
// panic: panic again
// panic: panic again and again

// 程序多次调用 panic 也不会影响 defer 函数的正常执行，所以使用 defer 进行收尾工作一般来说都是安全的。
