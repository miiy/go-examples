// 2.下面代码最后一行输出什么？请说明原因。
package main

import "fmt"

func main() {
	x := 1
	fmt.Println(x)
	{
		fmt.Println(x)
		i, x := 2, 2
		fmt.Println(i, x)
	}
	fmt.Println(x) // print ?
}

//参考答案及解析：输出1。知识点：变量隐藏。使用变量简短声明符号 := 时，如果符号左边有多个变量，只需要保证至少有一个变量是新声明的，并对已定义的变量尽进行赋值操作。但如果出现作用域之后，就会导致变量隐藏的问题，就像这个例子一样。
//
//这个坑很容易挖，但又很难发现。即使对于经验丰富的 Go 开发者而言，这也是一个非常常见的陷阱。
//
//引自：http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/
