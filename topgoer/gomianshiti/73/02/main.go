//2.下面的代码能编译通过吗？可以的话输出什么，请说明？
package main

var f = func(i int) {
	print("x")
}

func main() {
	f := func(i int) {
		print(i)
		if i > 0 {
			f(i - 1)
		}
	}
	f(10)
}

//参考答案及解析：10x。这道题一眼看上去会输出 109876543210，其实这是错误的答案，这里不是递归。假设 main() 函数里为 f2()，外面的为 f1()，当声明 f2() 时，调用的是已经完成声明的 f1()。
//
//看下面这段代码你应该会更容易理解一点：
//
//var x = 23
//
//func main() {
//	x := 2*x - 4
//	println(x)    // 输出:42
//}
