package main

import "fmt"

func counter() func() int {
	var i int

	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println(counter()()) // 1
	fmt.Println(counter()()) // 1
	f := counter()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
}

// 闭包的概念
// 闭包就是能够读取其他函数内部变量的函数。例如在javascript中，只有函数内部的子函数才能读取局部变量，所以闭包可以理解成“定义在一个函数内部的函数“。在本质上，闭包是将函数内部和函数外部连接起来的桥梁。
//
// https://baike.baidu.com/item/%E9%97%AD%E5%8C%85/10908873
// https://zhuanlan.zhihu.com/p/174619770
