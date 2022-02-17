//2.下面代码输出什么，请说明？
package main

import "fmt"

func main() {
	defer func() {
		fmt.Print(recover())
	}()
	defer func() {
		defer fmt.Print(recover())
		panic(1)
	}()
	defer recover()
	panic(2)
}

//参考答案及解析：21。recover() 必须在 defer() 函数中调用才有效，所以第 9 行代码捕获是无效的。在调用 defer() 时，便会计算函数的参数并压入栈中，所以执行第 6 行代码时，此时便会捕获 panic(2)；此后的 panic(1)，会被上一层的 recover() 捕获。所以输出 21。
//
//引自《Go语言101》
