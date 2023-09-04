// 1.下面两段代码输出什么。
package main

import "fmt"

// 1.
func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

// 2.
func main() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
}

//
//两段代码分别输出：
//
//[0 0 0 0 0 1 2 3]
//[1 2 3 4]
//
//参考解析：这道题考的是使用 append 向 slice 添加元素，第一段代码常见的错误是 [1 2 3]，需要注意。
