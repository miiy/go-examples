package main

import (
	"fmt"
	"runtime"
	"time"
)

func main()  {
	start := time.Now()

	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	chs := make([]chan int, cpus)
	for i := 0; i < len(chs); i++ {
		chs[i] = make(chan int, 1)
		go sum(i, chs[i])
	}
	sum := 0
	for _, ch := range chs {
		res := <- ch
		sum += res
	}
	end := time.Since(start).Seconds()
	fmt.Printf("%d %f", sum, end)
}

func sum(seq int, ch chan int)  {
	defer close(ch)
	sum := 0
	for i := 0; i < 10000000; i++ {
		sum += i
	}
	fmt.Println(seq)
	ch <- sum
}