//1.下面的代码输出什么？
package main

import "fmt"

type T struct {
	n int
}

func main() {
	ts := [2]T{}
	for i, t := range ts {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Print(ts)
}

//参考答案及解析：0 [{0} {9}]。知识点：for-range 循环数组。此时使用的是数组 ts 的副本，所以 t.n = 3 的赋值操作不会影响原数组。
