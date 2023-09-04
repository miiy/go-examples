package main

import "fmt"

// #4—Captured in a Loop

// 延迟函数在运行时将看到周围变量的最新值 -- 除了传递给它的参数。

// Example
// 在循环中 defer 闭包
func main() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

// Output:
// 3
// 3
// 3

// Why?
// 当代码运行时，deferred func 看到i最近的值。
// 这是因为，当 defer 被注册时，Go runtime 捕获 i 的地址，循环结束后 i 会变成 3。
// 所以，当 defer 运行，他们都将看起来相同 i，即 3.
//
//func ()  {
//	for i := 0; i < 3; i++ {
//		defer func() {
//			fmt.Println(i) // captures i address:0x2340fx00
//		}()
//	} // loop ends. i = 3
//	//
//	// fmt.Println(3)
//	// fmt.Println(3)
//	// fmt.Println(3)
//}

// Solution #1:
// 将参数直接传递给defer
//
//for i := 0; i < 3; i++ {
//	defer func(i int) {
//		fmt.Println(i)
//	}(i)
//}
//
// Output:
// 2
// 1
// 0
//
// Why it works?
// 因为这里,Go runtime 会创建不同的 i 变量，并使用正确的值将他们保存在一边。现在 defer 不在看到循环的原始 i 变量，它们看到自己的 i 局部变量。

// Solution #2:
// 这里，我们有意在for循环的作用域内使用新的 i 覆盖 i.我不喜欢这种风格。
//
//for i := 0; i < 3; i++ {
//	i := i
//	defer func() {
//		fmt.Println(i)
//	}()
//}

// Solution #3:
// 如果只有一个func调用，你可以直接使用它
//
//for i:=0; i < 3; i++ {
//	defer fmt.Println(i)
//}
