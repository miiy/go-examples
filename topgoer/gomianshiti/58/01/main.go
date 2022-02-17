//1.下面的代码输出什么？
package main

import "fmt"

type T struct {
	x int
	y *int
}

func main() {

	i := 20
	t := T{10,&i}

	p := &t.x

	*p++
	*p--

	t.y = p

	fmt.Println(*t.y)
}

//参考答案及解析：10。知识点：运算符优先级。如下规则：递增运算符 ++ 和递减运算符 – 的优先级低于解引用运算符 * 和取址运算符 &，解引用运算符和取址运算符的优先级低于选择器 . 中的属性选择操作符。
