//1.下面的代码输出什么？
package main

import "fmt"

type N int

func (n *N) test(){
	fmt.Println(*n)
}

func main()  {
	var n N = 10
	p := &n

	n++
	f1 := n.test

	n++
	f2 := p.test

	n++
	fmt.Println(n)

	f1()
	f2()
}

//参考答案及解析：13 13 13。知识点：方法值。当目标方法的接收者是指针类型时，那么被复制的就是指针。
//
//引自：《Go语言学习笔记》· 方法
