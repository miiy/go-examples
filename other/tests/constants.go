package main

import "fmt"

func main() {
	// 常量的定义格式：const identifier [type] = value
	const PI = 3.14159
	fmt.Println(PI)

	const STR1 string = "abc" // 显示定义
	const STR2 = "def" // 隐式定义 编译器可以根据变量的值来推断其类型
	fmt.Println(STR1, STR2)

	// 常量的值必须在编译时就能够确定
	// const NUM1 = getNum()

	const A, B = "a", "b"
	fmt.Println(A, B)

	// 常量还可以用作枚举
	const (
		UNKNOWN = 0
		FEMALE = 1
		MALE = 2
	)
	fmt.Println(UNKNOWN, FEMALE, MALE)
}

func getNum() int {
	return 1
}
