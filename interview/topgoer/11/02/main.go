//2.下面这段代码输出什么？
package main

import "fmt"

func main() {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}

//A. nil
//B. not nil
//C. compilation error
//
//参考答案及解析：A。当且仅当接口的动态值和动态类型都为 nil 时，接口类型值才为 nil。
