package main

import "fmt"

// 作用域
// 向 defer 关键字传入的函数会在函数返回前运行。
func main2() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// Output:
// 4
// 3
// 2
// 1
// 0

// 上述代码会倒序执行传入 defer 关键子的所有表达式，因为最后一次调用 defer 时传入了 fmt.Println(4)，所以这段代码会优先打印 4。

func main() {
	{
		defer fmt.Println("defer runs")
		fmt.Println("block ends")
	}

	fmt.Println("main ends")
}

// Output:
// block ends
// main ends
// defer runs

// defer传入的函数不是在推出代码块的作用域时执行的，他只会在当前函数和方法返回之前被调用。
