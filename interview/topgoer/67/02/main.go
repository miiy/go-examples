//2.下面的代码输出什么？
package main

import "fmt"

type T struct {
	n int
}

func main() {
	ts := [2]T{}
	for i := range ts[:] {
		switch t := &ts[i]; i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Print(ts)
}

//参考答案及解析：9 [{3} {9}]。知识点：for-range 切片。参考前几道题的解析，这道题的答案应该很明显。
//
//引自《Go语言101》
