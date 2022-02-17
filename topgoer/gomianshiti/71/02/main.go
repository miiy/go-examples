//2.下面代码输出什么，请说明原因。
package main

import "fmt"

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}
func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}
func main() {
	s := NewSlice()
	defer s.Add(1).Add(2)
	s.Add(3)
}

//参考答案及解析：132。这一题有两点需要注意：1.Add() 方法的返回值依然是指针类型 *Slice，所以可以循环调用方法 Add()；2.defer 函数的参数（包括接收者）是在 defer 语句出现的位置做计算的，而不是在函数执行的时候计算的，所以 s.Add(1) 会先于 s.Add(3) 执行。
