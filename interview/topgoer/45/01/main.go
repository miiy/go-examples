//1.下面代码有什么错误？
package main

func main() {
	one := 0
	one := 1
}

//参考答案及解析：变量重复声明。不能在单独的声明中重复声明一个变量，但在多变量声明的时候是可以的，但必须保证至少有一个变量是新声明的。
//
//修复代码：
//
//func main() {
//	one := 0
//	one, two := 1,2
//	one,two = two,one
//}
//
//引自：http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/