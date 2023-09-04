package main

import (
	"errors"
	"fmt"
)

// #5—panic/recover can get and return any type
// #5—panic/recover 可以获取和返回任何类型

// With string:
func errorly() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("error run run")
}

// output:
// error run run

func main() {
	errorly()
	//errorly2()
	//errorly3()
}

// With error:
func errorly2() {
	defer func() {
		fmt.Println(recover())
	}()
	panic(errors.New("error run run"))
}

// output:
// error run run

// Accepts any type
type myerror struct{}

func (myerror) String() string {
	return "myerror there!"
}

func errorly3() {
	defer func() {
		fmt.Println(recover())
	}()

	panic(myerror{})
}

// Why
// 因为 panic 接受一个 any 类型的参数
//
// recover 和 panic 在 go 中的声明如下：
//
// func panic(v any)
// func recover() any
//
// 基本上，它的工作原理是这样的：
// panic(value) -> recover() -> value
// recover 只是返回传递给 panic 的值
