//1.下面代码输出什么？
package main

import "fmt"

func main() {
	str := "hello"
	str[0] = 'x'
	fmt.Println(str)
}

//A. hello
//B. xello
//C. compilation error
//
//参考代码及解析：C。知识点：常量，Go 语言中的字符串是只读的。
