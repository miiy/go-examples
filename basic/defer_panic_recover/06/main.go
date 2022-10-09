package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		panic("error: panic in goroutine")
	}()
	fmt.Println("in main")
	time.Sleep(1 * time.Second)
}
