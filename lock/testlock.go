package grpcbackup

import (
	"fmt"
	"sync"
	"time"
)

func main2()  {
	testSyncLock()
}

func testWaitGroup() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("done.")
}

// % go run testlock.go |sort|uniq |wc -l
//      70
func testSyncLock() {
	lock := sync.Mutex{}
	a := 0
	for i := 1; i <= 1000; i++ {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			a = a + 1
			fmt.Println(a)
		}()
	}
	time.Sleep(time.Second)
}
