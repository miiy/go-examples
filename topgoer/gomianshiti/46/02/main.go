//2.下面代码输出什么？
package main

import "fmt"

const (
	x uint16 = 120
	y
	s = "abc"
	z
)

func main() {
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)
}

//参考答案及解析：知识点：常量。
//
//输出：
//
//uint16 120
//string abc
//
//常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
