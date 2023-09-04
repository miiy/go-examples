//3.下面这段代码输出什么？
package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}
func main() {
	v := 1
	incr(&v)
	fmt.Println(v)
}

//参考答案及解析：2。知识点：指针。p 是指针变量，指向变量 v，*p++操作的意思是取出变量 v 的值并执行加一操作，所以 v 的最终值是 2。
