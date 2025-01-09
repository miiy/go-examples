package main

import (
	"fmt"
	"time"
)

func main()  {
	c := make(chan int)
	go func() {
		time.Sleep(15 * 1e9)
		x := <- c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
}

/*

通道的阻塞性
通道的发送/接收操作在对方准备好之前是阻塞的

sending 10
sent 10
received 10
*/
