package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("In main()")
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main()")
}

func longWait()  {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) // sleep for 5 second
	fmt.Println("End of longWait()")
}

func shortWait()  {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 second
	fmt.Println("End of shortWait()")
}

/*

In main()
About to sleep in main()
Beginning shortWait()
Beginning longWait()
End of shortWait()
End of longWait()
At the end of main()

*/

/*
main(), longWait(), shortWait() 三个函数作为独立处理单元按顺序启动，然后并行开始运行
1e9 表示 1乘10的9次方，e=指数
如果不在main()中等待，协程会随着程序的结束而消亡
*/

/*
去掉go关键字
注释掉主线程中sleep
*/