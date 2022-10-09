package main

import (
	"fmt"
	"time"
)

func main() {
	case1()
	case2()
}

func case1() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1 * time.Second)
}

func case2() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1 * time.Second)
}

// case1 输出：
// 5
// 5
// 5
// 5
// 5

// case2 乱序输出0-4：
// 第一次运行
// 1
// 2
// 0
// 4
// 3

// 第二次运行
// 1
// 2
// 4
// 0
// 3
