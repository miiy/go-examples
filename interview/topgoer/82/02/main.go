//2.下面代码输出什么？
package main

import "fmt"

const (
	azero = iota
	aone  = iota
)

const (
	info  = "msg"
	bzero = iota
	bone  = iota
)

func main() {
	fmt.Println(azero, aone)
	fmt.Println(bzero, bone)
}

// 参考答案及解析：0 1 1 2。知识点：iota 的使用。这道题易错点在 bzero、bone 的值，在一个常量声明代码块中，如果 iota 没出现在第一行，则常量的初始值就是非 0 值。
//

