package main

import (
	"fmt"
	"sync"
)

func main()  {
	testWaitGroup()
}

func testWaitGroup() {
	var arr []int
	var mu sync.Mutex
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			mu.Lock()
			arr = append(arr, i)
			mu.Unlock()
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("%d done.\n", len(arr))
	fmt.Println(arr)
}