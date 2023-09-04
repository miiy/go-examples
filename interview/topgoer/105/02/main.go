// 2.下面代码输出什么？请简要说明。
package main

import (
	"fmt"
	"sync"
)

type MyMutex struct {
	count int
	sync.Mutex
}

func main()  {
	var mu MyMutex
	mu.Lock()
	var mu1 = mu
	mu.count++
	mu.Unlock()
	mu1.Lock()
	mu1.count++
	mu1.Unlock()
	fmt.Println(mu.count, mu1.count)
}

//A. 不能编译；
//B. 输出 1, 1；
//C. 输出 1, 2；
//D. fatal error；
//
//参考答案及解析：D。加锁后复制变量，会将锁的状态也复制，所以 mu1 其实是已经加锁状态，再加锁会死锁。

