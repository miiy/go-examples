//1.下面的代码能否正确输出？
package main

func main() {
	var fn1 = func() {}
	var fn2 = func() {}
	if fn1 != fn2 {
		println("fn1 not equal fn2")
	}
}

//参考答案及解析：编译错误
//
//invalid operation: fn1 != fn2 (func can only be compared to nil)
//
//函数只能与 nil 比较。