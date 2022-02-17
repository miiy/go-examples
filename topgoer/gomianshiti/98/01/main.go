//1.下面代码输出什么？
package main

import "fmt"

func main()  {
	a := 1
	for i := 0; i < 5; i++ {
		a := a + 1
		a = a *2
	}
	fmt.Println(a)
}

//参考答案及解析：1。知识点：变量的作用域。注意 for 语句的变量 a 是重新声明，它的作用范围只在 for 语句范围内。
