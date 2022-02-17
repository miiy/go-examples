//1.请指出下面代码的错误？
package main

import "fmt"

var gvar int

func main() {
	var one int
	two := 2
	var three int
	three = 3

	func(unused string) {
		fmt.Println("Unused arg. No compile error")
	}("what?")
}

//参考答案及解析：变量 one、two 和 three 声明未使用。知识点：未使用变量。如果有未使用的变量代码将编译失败。但也有例外，函数中声明的变量必须要使用，但可以有未使用的全局变量。函数的参数未使用也是可以的。
//
//如果你给未使用的变量分配了一个新值，代码也还是会编译失败。你需要在某个地方使用这个变量，才能让编译器愉快的编译。
//
//修复代码：
//
//func main() {
//	var one int
//	_ = one
//
//	two := 2
//	fmt.Println(two)
//
//	var three int
//	three = 3
//	one = three
//
//	var four int
//	four = four
//}
//
//另一个选择是注释掉或者移除未使用的变量 。
//
//引自：http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/
