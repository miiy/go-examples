//2.下面代码里的 counter 的输出值？
package main

import "fmt"

func main() {

	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
}

//A. 2
//B. 3
//C. 2 或 3
//
//参考答案及解析：C。for range map 是无序的，如果第一次循环到 A，则输出 3；否则输出 2。
