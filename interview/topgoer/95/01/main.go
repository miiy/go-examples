//1.下面代码输出什么？请简要说明。
package main

import "fmt"

type foo struct { Val int }
type bar struct { Val int }

func main()  {
	a := &foo{Val: 5}
	b := &foo{Val: 5}
	c := foo{Val: 5}
	d := bar{Val: 5}
	e := bar{Val: 5}
	f := bar{Val: 5}
	fmt.Print(a == b, c == foo(d), e == f)
}

//参考答案及解析：false true true。这道题唯一有疑问的地方就在第一个比较，Go 语言里没有引用变量，每个变量都占用一个惟一的内存位置，所以第一个比较输出 false。这个知识点在《Go 语言没有引用传递》有介绍。
