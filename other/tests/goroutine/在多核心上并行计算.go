package main

import "runtime"


func main() {
	DoAll()
}

func DoAll()  {
	ncpu := runtime.NumCPU()
	runtime.GOMAXPROCS(ncpu)

	ch := make(chan int, ncpu)
	for i := 0; i < ncpu; i ++ {
		go DoPart(ch)
	}
}

func DoPart(ch chan int)  {
	ch <- 1
}