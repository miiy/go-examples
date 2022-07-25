package main

import "sync"

var l sync.Mutex

func main() {
	total := 0

	for i := 0; i <= 10; i++ {
		l.Lock()
		go func() {
			total += 1
			l.Unlock()
		}()
	}
}
