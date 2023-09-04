//2.下面这段代码输出什么以及原因？
package main

import "fmt"

func hello() []string {
	return nil
}

func main() {
	h := hello
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

//A. nil
//B. not nil
//C. compilation error
//
//答案及解析：B。这道题目里面，是将 hello() 赋值给变量 h，而不是函数的返回值，所以输出 not nil。
