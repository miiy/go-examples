//1.下面代码输出什么？
package main

import "fmt"

func main()  {
	fmt.Println(doubleScore(0))
	fmt.Println(doubleScore(20.0))
	fmt.Println(doubleScore(50.0))
}

func doubleScore(source float32) (score float32) {
	defer func() {
		if score < 1 || score >= 100 {
			score = source
		}
	}()
	return source * 2
}

//参考答案及解析：输出 0 40 50。知识点：defer 语句与返回值。函数的 return value 不是原子操作，而是在编译器中分解为两部分：返回值赋值 和 return。