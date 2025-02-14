package main

import (
	"fmt"
	"time"
)

func main() {
	v := make(chan struct{})
	go func() {
		time.Sleep(2 * time.Second)
		v <- struct{}{}
	}()
	<-v
	fmt.Println(11)
}
