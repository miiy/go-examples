//2.下面的代码输出什么，请说明？
package main

import "fmt"

func main() {
	defer func() {
		fmt.Print(recover())
	}()
	defer func() {
		defer func() {
			fmt.Print(recover())
		}()
		panic(1)
	}()
	defer recover()
	panic(2)
}

//参考答案及解析：12。相关知识点请看 第64天 题目解析。
