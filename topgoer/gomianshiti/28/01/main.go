//1.下面的代码有什么问题？
package main

import "fmt"

func main() {
	fmt.Println([...]int{1} == [2]int{1})
	fmt.Println([]int{1} == []int{1})
}

//参考答案及解析：有两处错误
//
//go 中不同类型是不能比较的，而数组长度是数组类型的一部分，所以 […]int{1} 和 [2]int{1} 是两种不同的类型，不能比较；
//切片是不能比较的；
