//2.下面代码输出什么？
package main

import "fmt"

func main() {
	x := []string{"a", "b", "c"}
	for v := range x {
		fmt.Print(v)
	}
}

//参考答案及解析：012。注意区别下面代码段：
//
//func main() {
//	x := []string{"a", "b", "c"}
//	for _, v := range x {
//		fmt.Print(v)     //输出 abc
//	}
//}
