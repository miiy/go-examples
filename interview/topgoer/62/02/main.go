//2.下面代码输出什么？
package main

import "fmt"

func main() {
	var x int8 = -128
	var y = x/-1
	fmt.Println(y)
}

//参考答案及解析：-128。因为溢出。
