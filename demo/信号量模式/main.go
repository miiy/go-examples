package main

import (
	"fmt"
	"runtime"
	"time"
)

// https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.2.md#1427-信号量模式
// 信号量模式：协程通过在通道 ch 中放置一个值来处理结束的信号。main 协程等待 <-ch 直到从中获取到值。
func main() {
	start := time.Now()

	// demo1()
	// demo2()
	// demo3()
	demo4()

	fmt.Println("Time:", time.Since(start).String())
}

func demo1() {
	ch := make(chan int)
	go compute(ch)

	doSomethingElseForAWhile()

	result := <-ch // main 协程等待 <-ch 直到从中获取到值。
	fmt.Println(result)
}

// 这个信号也可以是其他的，不返回结果，比如下面这个协程中的匿名函数（lambda）协程：
func demo2() {
	ch := make(chan int)

	go func(ch chan int) {

		time.Sleep(time.Second * 2)

		ch <- 1 // Send a signal; value does not matter
	}(ch)

	time.Sleep(time.Second * 1)

	<-ch // Wait for goroutine to finish; discard sent value.
	fmt.Println("done.")
}

// 等待两个协程完成
func demo3() {
	done := make(chan bool)

	doSort := func(done chan bool) {
		time.Sleep(time.Second * 2)
		done <- true
	}

	go doSort(done)
	go doSort(done)
	t1 := <-done
	t2 := <-done

	fmt.Println(t1, t2)
}

// 缓存通道
func demo4() {
	ncpu := runtime.NumCPU()
	ch := make(chan bool, ncpu)
	doSomething := func(ch chan bool) {
		time.Sleep(time.Second * 2)
		ch <- true
	}

	for i := 0; i < ncpu; i++ {
		go doSomething(ch)
	}

	for i := 0; i < ncpu; i++ {
		done := <-ch
		fmt.Println(done)
	}

}

func compute(ch chan int) {

	rst := someComputation()

	ch <- rst // 在通道中放一个值来处理结束的信号
}

func someComputation() int {
	time.Sleep(time.Second * 2)
	return 100
}

func doSomethingElseForAWhile() {
	time.Sleep(time.Second * 1)
}
