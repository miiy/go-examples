//2.下面的代码有什么问题？
package main

import "fmt"

type T struct {
	n int
}

func getT() T {
	return T{}
}

func main() {
	getT().n = 1
}

//参考答案及解析：编译错误：
//
//cannot assign to getT().n
//
//直接返回的 T{} 无法寻址，不可直接赋值。
//
//修复代码：
//
//type T struct {
//	n int
//}
//
//func getT() T {
//	return T{}
//}
//
//func main() {
//	t := getT()
//	p := &t.n    // <=> p = &(t.n)
//	*p = 1
//	fmt.Println(t.n)
//}
