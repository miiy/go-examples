//1.下面这段代码输出什么？
package main

import "fmt"

type T struct {
	ls []int
}
func foo(t T) {
	t.ls[0] = 100
}
func main() {
	var t = T{
		ls: []int{1, 2, 3},
	}
	foo(t)
	fmt.Println(t.ls[0])
}

//A. 1
//B. 100
//C. compilation error
//
//参考答案及解析：B。调用 foo() 函数时虽然是传值，但 foo() 函数中，字段 ls 依旧可以看成是指向底层数组的指针。
