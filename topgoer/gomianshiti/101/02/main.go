//2.下面代码的功能是从小到大找出 17 和 38 的 3 个公倍数，请问下面的代码有什么问题？
package main

import (
	"fmt"
	"time"
)

var ch chan int = make(chan int)

func generate()  {
	for i := 17; i < 5000; i += 17 {
		ch <- i
		time.Sleep(1 * time.Millisecond)
	}
	close(ch)
}

func main()  {
	timeout := time.After(800 * time.Millisecond)
	go generate()
	found := 0
	for {
		select {
		case i, ok := <- ch:
			if ok {
				if i % 38 == 0 {
					fmt.Println(i, "is a multiple of 17 and 38")
					found++
					if found == 3 {
						break
					}
				}
			} else {
				break
			}
		case <- timeout:
			fmt.Println("timed out")
			break
		}
		fmt.Println("The end")
	}
}

//参考答案即解析：break 会跳出 select 块，但不会跳出 for 循环。这算是一个比较容易掉的坑。可以使用 break label 特性或者 goto 功能解决这个问题，这里使用 break label 作个示例。
//var ch chan int = make(chan int)
//
//func generate() {
//	for i := 17; i < 5000; i += 17 {
//		ch <- i
//		time.Sleep(1 * time.Millisecond)
//	}
//	close(ch)
//}
//
//func main()  {
//	timeout := time.After(800 * time.Millisecond)
//	go generate()
//	found := 0
//	MAIN_LOOP:
//	for {
//		select {
//		case i, ok := <- ch:
//			if ok {
//				if i % 38 == 0 {
//					fmt.Println(i, "is a multiple of 17 and 38")
//					found++
//					if found == 3 {
//						break MAIN_LOOP
//					}
//				}
//			} else {
//				break MAIN_LOOP
//			}
//		case <- timeout:
//			fmt.Println("timed out")
//			break MAIN_LOOP
//		}
//		fmt.Println("The end")
//	}
//}
