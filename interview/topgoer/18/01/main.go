//1.f1()、f2()、f3() 函数分别返回什么？
package main

import "fmt"

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}

// 1 5 1
