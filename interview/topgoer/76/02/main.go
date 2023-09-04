//2.下面的代码有什么问题？
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("1")
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()
}

//参考答案及解析：协程里面，使用 wg.Add(1) 但是没有 wg.Done()，导致 panic()。
