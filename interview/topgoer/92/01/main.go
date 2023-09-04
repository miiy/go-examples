//1.下面代码输出什么？
package main

import "fmt"

var x int

func init() {
	x++
}

func main() {
	init()
	fmt.Println(x)
}

//参考答案及解析：编译失败。init() 函数不能被其他函数调用，包括 main() 函数。
