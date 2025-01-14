package main

import "fmt"

func main()  {
	ch1 := make(chan int)
	go pump(ch1)       // pump hangs
	fmt.Println(<-ch1) // prints only 0
}

func pump(ch chan int)  {
	for i := 0; ; i++ {
		ch <- i
	}
}

/*
一个协程在无线循环中给通道发送整数数据，因为没有接受者，只输出了一个数字0
 */
