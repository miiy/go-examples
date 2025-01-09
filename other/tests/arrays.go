package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 数组声明 var identifier [len]type，Go语言中数组是值类型
	var arr1 [5]int
	fmt.Println(arr1)

	// 数组初始化
	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	arr2[0] = 2;
	fmt.Println(arr2)

	// 切片声明
	var s1 []int
	fmt.Println(s1)

	// 切片初始化
	var s2 []int = arr1[0:5]
	fmt.Println(s2)
	fmt.Println(reflect.TypeOf(arr1))
	fmt.Println(reflect.TypeOf(s1))

	// map 初始化 var map1 map[keytype]valuetype
	return

	// for
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	// for-range
	for i, v := range arr1 {
		fmt.Printf("key: %d, value: %d \n",i, v)
	}

	for i, _ := range arr1 {
		fmt.Printf("key: %d, value: %d \n",i, arr1[i])
	}

	for i := range arr1 {
		fmt.Printf("key: %d \n", i)
	}

	for range arr1 {
		fmt.Println()
	}
	//
	//arr2 := arr1
	//arr2[2] = 100
	//fmt.Println(arr1)
	//fmt.Println(arr2)

}
