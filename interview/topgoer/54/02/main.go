// 2.下面的代码输出什么？
package main

import "fmt"

type N int

func (n N) test() {
	fmt.Println(n)
}

func main() {
	var n N = 10
	fmt.Println(n)

	n++
	f1 := N.test
	f1(n)

	n++
	f2 := (*N).test
	f2(&n)

	say := (*A).Say
	say(new(A))
	new(A).Say()
}

type A struct {
}

func (a A) Say() {
	fmt.Println(111)
}

//参考答案及解析：10 11 12。知识点：方法表达式。通过类型引用的方法表达式会被还原成普通函数样式，接收者是第一个参数，调用时显示传参。类型可以是 T 或 *T，只要目标方法存在于该类型的方法集中就可以。
//
//还可以直接使用方法表达式调用：
//
//func main()  {
//	var n N = 10
//
//	fmt.Println(n)
//
//	n++
//	N.test(n)
//
//	n++
//	(*N).test(&n)
//}
//
//引自：《Go语言学习笔记》· 方法
