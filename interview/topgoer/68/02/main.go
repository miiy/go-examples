//2.下面代码输出什么，请说明。
package main

func main() {
	x := []int{0, 1, 2}
	y := [3]*int{}
	for i, v := range x {
		defer func() {
			print(v)
		}()
		y[i] = &v
	}
	print(*y[0], *y[1], *y[2])
}

//参考答案及解析：22222。知识点：defer()、for-range。for-range 虽然使用的是 :=，但是 v 不会重新声明，可以打印 v 的地址验证下。
