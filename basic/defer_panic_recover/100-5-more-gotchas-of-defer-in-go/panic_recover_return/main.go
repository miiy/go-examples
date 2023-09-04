package main

import "fmt"

func foo() (i int) {
	i = 0
	defer func() {
		fmt.Println(i)
	}()
	return 2
}

func main() {
	foo()
}

//
//output:
//2
//解释：名为 test 的函数一直运行至结束，然后 defer 函数会被执行且会因为值为 nil 而产生 panic 异常。然而值得注意的是，run() 的声明是没有问题，因为在test函数运行完成后它才会被调用。
