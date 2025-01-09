package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int

func main() {
	var lock = sync.Mutex{}
	for i := 0; i < 10; i++ {
		go add(i, &lock)
	}
	time.Sleep(time.Second * 2)
	fmt.Println(counter)
}

func add(i int, lock *sync.Mutex) {
	fmt.Println(i)
	lock.Lock()
	counter++
	lock.Unlock()
}
