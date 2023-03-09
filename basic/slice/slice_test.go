package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

//	slice底层数据结构
//
//	go/src/runtime/slice.go:15
//	type slice struct {
//		array unsafe.Pointer
//		len   int
//		cap   int
//	}
//
//	go/src/reflect/value.go:2758
//	type SliceHeader struct {
//		Data uintptr
//		Len  int
//		Cap  int
//	}

// TestChangeSlice 通过函数改变切片
func TestChangeSlice(t *testing.T) {
	s := []int{1, 2, 3}
	fmt.Println(s) // [1 2 3]

	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	dataPtr := unsafe.Pointer(hdr.Data)
	fmt.Printf("hdr: %+v, dataPtr: %p\n", hdr, dataPtr) // hdr: &{Data:824634933320 Len:3 Cap:3}, dataPtr: 0xc000128048

	changeSlice(s) // changeSlice(s) 参数传递的是一个 SliceHeader 结构体，Data指针指向的是一个数组
	fmt.Println(s) // [1 0 3]
}

func changeSlice(s []int) {
	s[1] = 0
}

// TestAppendSlice 切片扩容
// appendSlice 函数中 slice 发生了扩容，Data 指针指向了新的数组地址，Len，Cap 已经改变
func TestAppendSlice(t *testing.T) {
	s := []int{1, 2, 3}
	fmt.Println(s) // [1 2 3]

	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	dataPtr := unsafe.Pointer(hdr.Data)
	fmt.Printf("hdr: %+v, dataPtr: %p\n", hdr, dataPtr) // hdr: &{Data:824633835736 Len:3 Cap:3}, dataPtr: 0xc00001c0d8

	appendSlice(s)
	fmt.Println(s) // [1 2 3]

	hdr = (*reflect.SliceHeader)(unsafe.Pointer(&s))
	dataPtr = unsafe.Pointer(hdr.Data)
	fmt.Printf("hdr: %+v, dataPtr: %p\n", hdr, dataPtr) // hdr: &{Data:824633835736 Len:3 Cap:3}, dataPtr: 0xc00001c0d8

	// 获取底层的数组，切片扩容后，底层的数组已改变
	hdr = (*reflect.SliceHeader)(unsafe.Pointer(&s))
	data := *(*[4]int)(unsafe.Pointer(hdr.Data))
	fmt.Println(data) // [1 2 3 824634449992] 越界
}

func appendSlice(s []int) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	dataPtr := unsafe.Pointer(hdr.Data)
	fmt.Printf("hdr: %+v, dataPtr: %p\n", hdr, dataPtr) // hdr: &{Data:824633835736 Len:3 Cap:3}, dataPtr: 0xc00001c0d8

	s = append(s, 4)
	fmt.Println(s) // [1 2 3 4]

	hdr = (*reflect.SliceHeader)(unsafe.Pointer(&s))
	dataPtr = unsafe.Pointer(hdr.Data)
	fmt.Printf("hdr: %+v, dataPtr: %p\n", hdr, dataPtr) // hdr: &{Data:824633868624 Len:4 Cap:6}, dataPtr: 0xc000024150
}

// TestAppendSlice2 切片添加2
// 切片未发生扩容，Data指针指向的是同一数组，但是长度发生了改变
func TestAppendSlice2(t *testing.T) {
	s := make([]int, 3, 10) // 创建一个长度为3，容量为10的切片
	s = append(s, []int{1, 2, 3}...)
	fmt.Println(s) // [0 0 0 1 2 3]
	appendSlice2(s)
	fmt.Println(s) // [0 0 0 1 2 3]

	// 获取底层的数组，切片未扩容指向的是统一数组，只是长度变了
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	data := *(*[7]int)(unsafe.Pointer(hdr.Data))
	fmt.Println(data) // [0 0 0 1 2 3 4]
}

func appendSlice2(s []int) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	dataPtr := unsafe.Pointer(hdr.Data)
	fmt.Printf("hdr: %+v, dataPtr: %p\n", hdr, dataPtr) // hdr: &{Data:824633827888 Len:6 Cap:10}, dataPtr: 0xc00001a230

	s = append(s, 4)
	fmt.Println(s) // [0 0 0 1 2 3 4]

	hdr = (*reflect.SliceHeader)(unsafe.Pointer(&s))
	dataPtr = unsafe.Pointer(hdr.Data)
	fmt.Printf("hdr: %+v, dataPtr: %p\n", hdr, dataPtr) // hdr: &{Data:824633827888 Len:7 Cap:10}, dataPtr: 0xc00001a230
}
