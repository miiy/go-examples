//2.下面的代码输出什么？
package main

import "fmt"

func test(i int) (ret int) {
	ret = i *2
	if ret > 10 {
		ret := 10
		return
	}
	return
}

func main()  {
	result := test(10)
	fmt.Println(result)
}

//参考答案即解析：编译错误。知识点：变量的作用域。编译错误信息：ret is shadowed during return。
