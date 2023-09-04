//1.下面的代码有什么问题？
package main

import "fmt"

func main() {
	const x = 123
	const y = 1.23
	fmt.Println(x)
}

//参考答案及解析：编译可以通过。知识点：常量。常量是一个简单值的标识符，在程序运行时，不会被修改的量。不像变量，常量未使用是能编译通过的。
