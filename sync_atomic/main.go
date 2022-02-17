// https://zhuanlan.zhihu.com/p/210153612
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var value int64 = 0

func changeValue(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10000; i++ {
		value ++
	}
}

func changeValueAtomic(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10000 ;i++ {
		atomic.AddInt64(&value, 1)
	}
}

func demo1 () {
	fmt.Println(value)
	var wg sync.WaitGroup
	wg.Add(2)
	go changeValue(&wg)
	go changeValue(&wg)
	wg.Wait()
	fmt.Println(value)
}

func demo2 () {
	fmt.Println(value)
	var wg sync.WaitGroup
	wg.Add(2)
	go changeValueAtomic(&wg)
	go changeValueAtomic(&wg)
	wg.Wait()
	fmt.Println(value)
}

func main() {
	//demo1()
	demo2()
}
