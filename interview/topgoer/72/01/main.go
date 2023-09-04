//1.下面的代码输出什么，请说明。
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
	defer func() {
		s.Add(1).Add(2)
	}()
	s.Add(3)
}

//参考答案及解析：312。对比昨天的第二题，本题的 s.Add(1).Add(2) 作为一个整体包在一个匿名函数中，会延迟执行。
