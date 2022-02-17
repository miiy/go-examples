// 2.下面的代码有什么问题？
package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	sync.Mutex
}

func (d data) test(s string)  {
	d.Lock()
	defer d.Unlock()

	for i:=0;i<5 ;i++  {
		fmt.Println(s,i)
		time.Sleep(time.Second)
	}
}


func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	var d data

	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}

//参考答案及解析：锁失效。将 Mutex 作为匿名字段时，相关的方法必须使用指针接收者，否则会导致锁机制失效。
//
//修复代码：
//
//func (d *data) test(s string)  {     // 指针接收者
//	d.Lock()
//	defer d.Unlock()
//
//	for i:=0;i<5 ;i++  {
//		fmt.Println(s,i)
//		time.Sleep(time.Second)
//	}
//}
//
//或者可以通过嵌入 *Mutex 来避免复制的问题，但需要初始化。
//
//type data struct {
//	*sync.Mutex     // *Mutex
//}
//
//func (d data) test(s string) {    // 值方法
//	d.Lock()
//	defer d.Unlock()
//
//	for i := 0; i < 5; i++ {
//		fmt.Println(s, i)
//		time.Sleep(time.Second)
//	}
//}
//
//func main() {
//
//	var wg sync.WaitGroup
//	wg.Add(2)
//
//	d := data{new(sync.Mutex)}   // 初始化
//
//	go func() {
//		defer wg.Done()
//		d.test("read")
//	}()
//
//	go func() {
//		defer wg.Done()
//		d.test("write")
//	}()
//
//	wg.Wait()
//}
//
//引自：《Go 语言学习笔记》· 同步
