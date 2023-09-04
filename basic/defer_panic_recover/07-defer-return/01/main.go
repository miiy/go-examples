package main

import "fmt"

// 无名返回值
func main() {
	fmt.Println("return:", test())
}

func test() int {
	var i int
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

// output:
// defer1: 1
// defer2: 2
// 0

// return 非原子性
// go 语言中函数的 return 不是原子操作，在底层分为两步来执行
// 1. 返回值赋值
// 2. 真正的RET返回

// 函数中如果存在 defer，那么执行的时机是在第一步和第二步之间
// 1. 返回值赋值
// 2. defer
// 3. 真正的RET返回

// 返回值没有命名，所以 return 之前
// 1. 创建临时变量作为返回值 r := i
// 2. 执行 defer
// 3. 返回 return r

func test2() int {
	t := 5
	defer func() {
		t += 5
	}()
	return t
}

// Output:
// return: 5
