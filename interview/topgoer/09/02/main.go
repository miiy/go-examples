//2.下面这段代码输出什么？
package main

import "fmt"

type person struct {
	name string
}

func main() {
	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p])
}

//A.0
//B.1
//C.Compilation error
//
//参考答案及解析：A。打印一个 map 中不存在的值时，返回元素类型的零值。这个例子中，m 的类型是 map[person]int，因为 m 中不存在 p，所以打印 int 类型的零值，即 0。
