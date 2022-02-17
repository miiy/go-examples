//2.下面的代码输出什么？请简要说明。
package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var chain string

func main()  {
	chain = "main"
	A()
	fmt.Print(chain)
}

func A()  {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"
	B()
}

func B()  {
	chain = chain + " --> B"
	C()
}

func C()  {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> C"
}

//A. 不能编译；
//B. 输出 main –> A –> B –> C；
//C. 输出 main；
//D. fatal error；
//
//参考答案即解析：D。使用 Lock() 加锁后，不能再继续对其加锁，直到利用 Unlock() 解锁后才能再加锁。