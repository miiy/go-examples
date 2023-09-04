//1.下面代码能编译通过吗？
package main

import "fmt"

func main()  {
	true := false
	fmt.Println(true)
}

//参考答案即解析：编译通过。true 是预定义标识符可以用作变量名，但是不建议这么做。
