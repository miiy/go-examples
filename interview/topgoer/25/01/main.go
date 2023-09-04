//1.下面这段代码输出什么？为什么？
package main

import "fmt"

func (i int) PrintInt ()  {
	fmt.Println(i)
}

func main() {
	var i int = 1
	i.PrintInt()
}

//A. 1
//B. compilation error
//
//参考答案及解析：B。基于类型创建的方法必须定义在同一个包内，上面的代码基于 int 类型创建了 PrintInt() 方法，由于 int 类型和方法 PrintInt() 定义在不同的包内，所以编译出错。
//
//解决的办法可以定义一种新的类型：
//
//type Myint int
//
//func (i Myint) PrintInt ()  {
//	fmt.Println(i)
//}
//
//func main() {
//	var i Myint = 1
//	i.PrintInt()
//}
