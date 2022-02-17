//2.下面的代码有什么问题？
package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		go func(wg sync.WaitGroup, i int) {
			wg.Add(1)
			fmt.Printf("i:%d\n", i)
			wg.Done()
		}(wg, i)
	}

	wg.Wait()

	fmt.Println("exit")
}

//参考答案及解析：知识点：WaitGroup 的使用。存在两个问题：
//
//在协程中使用 wg.Add()；
//使用了 sync.WaitGroup 副本；
//
//修复代码：
//
//func main() {
//
//	wg := sync.WaitGroup{}
//
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func(i int) {
//			fmt.Printf("i:%d\n", i)
//			wg.Done()
//		}(i)
//	}
//
//	wg.Wait()
//
//	fmt.Println("exit")
//}
//
//// 或者
//
//func main() {
//
//	wg := &sync.WaitGroup{}
//
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func(wg *sync.WaitGroup,i int) {
//			fmt.Printf("i:%d\n", i)
//			wg.Done()
//		}(wg,i)
//	}
//
//	wg.Wait()
//
//	fmt.Println("exit")
//}
