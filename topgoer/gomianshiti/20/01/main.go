//1.下面这段代码正确的输出是什么？
package main

import "fmt"

func f() {
	defer fmt.Println("D")
	fmt.Println("F")
}

func main() {
	f()
	fmt.Println("M")
}

//A. F M D
//B. D F M
//C. F D M
//
//参考答案及解析：C。被调用函数里的 defer 语句在返回之前就会被执行，所以输出顺序是 F D M。
