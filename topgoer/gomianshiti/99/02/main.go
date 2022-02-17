package main

import "fmt"

func watShadowDefer(i int) (ret int) {
	ret = i * 2
	if ret > 10 {
		ret := 10
		defer func() {
			ret = ret + 1
		}()
	}
	return
}

func main()  {
	result := watShadowDefer(50)
	fmt.Println(result)
}

// 参考答案即解析：100。知识点：变量作用域和defer 返回值。
