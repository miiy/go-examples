package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++ // 修改是x不是返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值=x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ // 修改的是x
	}()
	return x // 1. r=5, x++, return r
}

func f4() (x int) {
	defer func(x int) {
		x++ // 改变的是函数中x的副本
	}(x)
	return 5 // 返回值 = x = 5
}

func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5 // 1. x= 5, 2. x++, 3. ret
}

func f6() (x int) {
	defer func(x *int) *int {
		(*x)++
		return x
	}(&x)
	return 5
}

func main() {
	fmt.Println("f1:", f1())
	fmt.Println("f2:", f2())
	fmt.Println("f3:", f3())
	fmt.Println("f4:", f4())
	fmt.Println("f5:", f5())
	fmt.Println("f6:", f6())
}

// Output:
// f1: 5
// f2: 6
// f3: 5
// f4: 5
// f5: 5
// f6: 6
