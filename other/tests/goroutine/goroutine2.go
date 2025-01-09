package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}

func sendData(ch chan string)  {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData(ch chan string)  {
	var input string
	// time.Sleep(2e9)
	for {
		input = <- ch
		fmt.Println(input)
	}
}

/*
通信操作符 <-

main() 函数启动了两个协程
	sendData()
	getData()
main()等待1秒让两个协程完成

*/