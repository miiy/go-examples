//2.下面代码输出什么？
package main

import "fmt"

type P *int
type Q *int

func main()  {
	var p P = new(int)
	*p += 8
	var x *int = p
	var q Q = x
	*q++
	fmt.Println(*p, *q)
}

//A.8 8
//B.8 9
//C.9 9
//
//参考答案及解析：C。指针变量指向相同的地址。
