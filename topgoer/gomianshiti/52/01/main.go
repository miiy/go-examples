//1.下面的代码有什么问题？
package main

type X struct {}

func (x *X) test()  {
	println(x)
}

func main() {
	var a *X
	a.test()
	X{}.test()
}

//参考答案及解析：X{} 是不可寻址的，不能直接调用方法。知识点：在方法中，指针类型的接收者必须是合法指针（包括 nil）,或能获取实例地址。
//
//修复代码：
//
//func main() {
//	var a *X
//	a.test()    // 相当于 test(nil)
//	var x = X{}
//	x.test()
//}
//
//引自：《Go语言学习笔记》· 方法