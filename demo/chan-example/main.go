// 在多核心上并行计算
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU()) // GO 1.5开始默认
	start := time.Now()

	doAll()

	end := time.Now()
	fmt.Println("Time:", end.Sub(start))
}

func doPart(seq int, sem chan int) {
	// do the part of the computation
	time.Sleep(time.Second * 2)
	fmt.Printf("Task %d finished.\n", seq)
	sem <- 1 // signal that this piece is done
}

func doAll() {
	ncpu := runtime.NumCPU()

	sem := make(chan int, ncpu) // Buffering optional but sensible
	for i := 0; i < ncpu; i++ {
		go doPart(i, sem)
	}

	// Drain the channel sem, waiting for NCPU tasks to complete
	for i := 0; i < ncpu; i++ {
		// <- sem  // wait for one task to complete
		done := <-sem
		fmt.Printf("Task: %d, %d\n", i, done)
	}
	// All done.
	fmt.Println("All done.")
}
