package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 8字节
	fmt.Println(unsafe.Sizeof(int(1)))
	// 4字节
	fmt.Println(unsafe.Sizeof(int32(1)))
	// 8字节
	i := int(0)
	p := &i
	fmt.Println(unsafe.Sizeof(p))
	// int ptr 大小跟随系统字长，取决于系统是32位的还是64位的
	//

	type k struct {
	}
	a := k{}
	b := int(0)
	c := k{}
	// 0 空结构体有地址没长度，空结构体的长度均相同，指针指向zerobase，zerobase是所有长度为0字节的地址
	fmt.Println(unsafe.Sizeof(a))
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c)

	// 16
	fmt.Println(unsafe.Sizeof("字符串"))
	// 16
	fmt.Println(unsafe.Sizeof("字符串字符串字符串字符串字符串字符串字符串"))
	// 字符串底层是
	// type stringStruct struct {
	//	  str unsafe.Pointer
	//    len int
	// }

	s := "字符串a"
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	// 10 Unicode每个字至少需要3个字节表示，UTF-8是Unicode的一种变长格式
	fmt.Println(sh.Len)

}
