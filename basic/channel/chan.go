package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan struct{}, 10)
	defer close(ch)
	for {
		ch <- struct{}{}
		go consume(ch)
	}

}

func consume(ch chan struct{}) {
	//
	time.Sleep(time.Second * 1)
	fmt.Println(runtime.NumGoroutine())

	<-ch
}
