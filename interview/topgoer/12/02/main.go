//2.下面这段代码输出什么？
package main

import "fmt"

func main() {
	i := -5
	j := +5
	fmt.Printf("%+d %+d", i, j)
}

//A. -5 +5
//B. +5 +5
//C. 0 0
//
//参考答案及解析：A。%d表示输出十进制数字，+表示输出数值的符号。这里不表示取反。
