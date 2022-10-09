package main

import (
	"fmt"
	"sync"
)

func main() {
	demo1()
}

func demo1() {
	var m sync.Map
	// Store 设置k, v
	m.Store("name", "zhang san")
	m.Store("name", "li si")
	m.Store("age", 20)

	// Range 遍历
	m.Range(func(k, v interface{}) bool {
		fmt.Printf("%s: %v\n", k, v)
		return true
	})

	fmt.Println("-------------------")

	// Delete 删除key
	m.Delete("age")

	// Load 返回map中key的值，如果不存在返回 nil, false
	name, ok := m.Load("name")
	fmt.Println(name, ok)
	age, ok := m.Load("age")
	fmt.Println(age, ok)

	fmt.Println("-------------------")

	m.Range(func(k, v interface{}) bool {
		fmt.Printf("%s: %v\n", k, v)
		return true
	})

	fmt.Println("-------------------")

	// LoadOrStore 如果存在键返回键的现有值
	// 如果不存在存储并返回值
	actual, loaded := m.LoadOrStore("phone", "123456")
	fmt.Println(actual, loaded)
	actual, loaded = m.LoadOrStore("phone", "654321")
	fmt.Println(actual, loaded)

	fmt.Println("-------------------")

	// LoadAndDelete 删除键的值
	value, loaded := m.LoadAndDelete("birthday")
	fmt.Println(value, loaded)

	value, loaded = m.LoadAndDelete("phone")
	fmt.Println(value, loaded)

	fmt.Println("-------------------")

	m.Range(func(k, v interface{}) bool {
		fmt.Printf("%s: %v\n", k, v)
		return true
	})
}

