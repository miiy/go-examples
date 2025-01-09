package main

import "fmt"

func main() {
	// if
	if 1 > 0 {
		fmt.Println(true)
	}

	// if else
	if 1 > 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	// if else if
	if 1 > 0 {
		fmt.Println(true)
	} else if 2 > 0{
		fmt.Println(false)
	} else {

	}

	// if 可以包含一个初始化语句（如：给一个变量赋值）。这种写法具有固定的格式（在初始化语句后方必须加上分号）：
	if a := 1; a > 0 {
		fmt.Println(true)
	}

	// switch
	// 不会自动的执行下一个分支的代码
	num1 := 100
	switch num1 {
	case 100:
		fmt.Println("100")
	default:
		fmt.Println("default")
	}

	switch num1 {
	case 100:
		fmt.Println("100")
		fallthrough
	case 101:
		fmt.Println("101")
		fallthrough
	default:
		fmt.Println("default")
	}

}




// switch

// select

// for

// break

// continue

// goto


