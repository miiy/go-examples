//2.下面代码是否能编译通过？如果通过，输出什么？
package main

import "fmt"

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

func main() {
	var x *int = nil
	Foo(x)
}

//参考答案及解析：non-empty interface 考点：interface 的内部结构，我们知道接口除了有静态类型，还有动态类型和动态值，当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。这里的 x 的动态类型是 *int，所以 x 不为 nil。
