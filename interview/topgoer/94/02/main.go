//2.下面哪些函数不能通过编译？
package main

import "fmt"

func A(string string) string {
	return string + string
}

func B(len int) int {
	return len + len
}

func C(val, default string) string {
	if val == "" {
		return default
	}
	return val
}

//参考答案及解析：C() 函数不能通过编译。C() 函数的 default 属于关键字。string 和 len 是预定义标识符，可以在局部使用。nil 也可以当做变量使用，不过不建议写这样的代码，可读性不好，小心被接手你代码的人胖揍。
//
//var nil = new(int)
//
//func main() {
//	var p *int
//	if p == nil {
//		fmt.Println("p is nil")
//	} else {
//		fmt.Println("p is not nil")
//	}
//}
//
