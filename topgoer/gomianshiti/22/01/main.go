//1.下面的代码有几处语法问题，各是什么？
package main

import (
"fmt"
)

func main() {
	var x string = nil
	if x == nil {
		x = "default"
	}
	fmt.Println(x)
}

//参考答案及解析：两个地方有语法问题。golang 的字符串类型是不能赋值 nil 的，也不能跟 nil 比较。
