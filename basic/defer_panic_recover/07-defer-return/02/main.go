package main

import "fmt"

// 有名返回值
func main() {
	fmt.Println("return:", demo())
}

func demo() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i
}

// Output:
// defer1: 1
// defer2: 2
// return: 2

// 先为返回值赋值，然后执行 defer，然后 return 到函数调用处。
// 1. 返回值已经定义不会产生历史变量 result = 0
// 2. 执行 defer result++
// 3. return result
