package main

import "fmt"

func main() {
	test1()
	test2()
	test3()
	test4()
}

func test1()  {
	// 声明变量 var identifier type

	var a int
	var b bool
	var str string

	//var (
	//	a int
	//	b bool
	//	str string
	//)

	fmt.Println(a, b, str)
}

func test2()  {
	// 声明并赋值 var identifier [type] = value
	var a int = 5
	var i = 5
	var b bool = false
	var str string = "hello"
	fmt.Println(a, i, b, str)

	// 自动判断类型
	var a1 = 5
	var b1 = false
	var str1 = "hello"

	//var (
	//	a1 = 5
	//	b1 = false
	//	str1 = "hello"
	//)

	fmt.Println(a1, b1, str1)

}

func test3()  {
	// 简短声明语法 := 类型由编译器自动推断
	a := 1
	b := false
	str := "hello"
	fmt.Println(a, b, str)
}

func test4()  {
	// 声明多个变量 var identifier1, identifier2 type
	// 同一类型的多个变量可以声明在同一行
	var a, b, c int
	fmt.Println(a, b, c)

	var d, e, f int = 1, 2, 3
	fmt.Println(d, e, f)

	g, h, i := 1, false, "hello"
	fmt.Println(g, h, i)
}
