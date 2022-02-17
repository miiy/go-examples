//2.下面的代码有什么问题？
package main

func main() {
	var x = nil
	_ = x
}

//参考答案及解析：nil 用于表示 interface、函数、maps、slices 和 channels 的“零值”。如果不指定变量的类型，编译器猜不出变量的具体类型，导致编译错误。
//
//修复代码：
//
//func main() {
//	var x interface{} = nil
//	_ = x
//}