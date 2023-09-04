//2.下面代码输出什么，为什么？
package main

import "fmt"

const (
	Century = 100
	Decade  = 010
	Year    = 001
)

func main() {
	fmt.Println(Century + 2*Decade + 2*Year)
}

//参考答案及解析：118。知识点：进制数。Go 语言里面，八进制数以 0 开头，十六进制数以 0x 开头，所以 Decade 表示十进制的 8。
