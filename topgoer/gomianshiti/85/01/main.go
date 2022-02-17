//1.下面这段代码输出什么？请简要说明。
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimRight("ABBA", "BA"))
}

//参考答案及解析：输出空字符。这是一个大多数人遇到的坑，TrimRight() 会将第二个参数字符串里面所有的字符拿出来处理，只要与其中任何一个字符相等，便会将其删除。想正确地截取字符串，可以参考 TrimSuffix() 函数。
