package main

import (
	"fmt"
	"sync"
)

func main() {
	number, letter := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-letter:
				if i >= 'Z' {
					wg.Done()
					return
				}
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				number <- true
			}
		}
	}(&wg)

	number <- true
	wg.Wait()
}

// 交替打印数字和字母
//
// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
//
// 解题思路
// 问题很简单，使用 channel 来控制打印的进度。使用两个 channel ，来分别控制数字和字母的打印序列， 数字打印完成后通过 channel 通知字母打印, 字母打印完成后通知数字打印，然后周而复始的工作。
//
// 源码解析
// 这里用到了两个channel负责通知，letter负责通知打印字母的goroutine来打印字母，number用来通知打印数字的goroutine打印数字。
// wait用来等待字母打印完成后退出循环。
