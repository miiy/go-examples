//2.下面代码输出什么？
package main

import "fmt"

func main() {
	var k = 9
	for k = range []int{} {}
	fmt.Println(k)

	for k = 0; k < 3; k++ {
	}
	fmt.Println(k)


	for k = range (*[3]int)(nil) {
	}
	fmt.Println(k)
}

//参考答案及解析：932。
