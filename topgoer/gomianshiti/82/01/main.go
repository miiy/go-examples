//1.下面这段代码输出什么？
package main

import "fmt"

func main() {
	count := 0
	for i := range [256]struct{}{} {
		m, n := byte(i), int8(i)
		if n == -n {
			count++
		}
		if m == -m {
			count++
		}
	}
	fmt.Println(count)
}

//参考答案及解析：4。知识点：数值溢出。当 i 的值为 0、128 是会发生相等情况，注意 byte 是 uint8 的别名。
//
//引自《Go语言101》
