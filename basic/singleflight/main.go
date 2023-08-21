package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
)

// https://pkg.go.dev/golang.org/x/sync/singleflight
// Package singleflight provides a duplicate function call suppression mechanism."
// singleflight 包提供了一个重复函数调用抑制机制。

func main() {
	g := new(singleflight.Group)

	block := make(chan struct{})
	res1c := g.DoChan("key", func() (interface{}, error) {
		<-block
		return "func 1", nil
	})
	res2c := g.DoChan("key", func() (interface{}, error) {
		<-block
		return "func 2", nil
	})
	close(block)
	res1 := <-res1c
	res2 := <-res2c

	// Results are shared by functions executed with duplicate keys.
	// 结果是由执行具有重复key的函数共享的。
	fmt.Println("Shared:", res2.Shared)
	// Only the first function is executed: it is registered and started with "key",
	// and doesn't complete before the second funtion is registered with a duplicate key.
	// 只有第一个函数被执行：它被注册并使用 "key" 开始，
	// 并且在第二个函数使用重复的 key 注册之前并没有完成。
	fmt.Println("Equal results:", res1.Val.(string) == res2.Val.(string))
	fmt.Println("Result:", res1.Val)
}
