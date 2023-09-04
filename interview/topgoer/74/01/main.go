//1.下面代码有什么问题，请说明？
package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		for i:=0;i<10 ;i++  {
			fmt.Println(i)
		}
	}()

	for {}
}

//参考答案及解析：for {} 独占 CPU 资源导致其他 Goroutine 饿死。

//可以通过阻塞的方式避免 CPU 占用，修复代码：

func main2() {
	runtime.GOMAXPROCS(1)

	go func() {
		for i:=0;i<10 ;i++  {
			fmt.Println(i)
		}
		os.Exit(0)
	}()

	select {}
}

//引自《Go语言高级编程》
