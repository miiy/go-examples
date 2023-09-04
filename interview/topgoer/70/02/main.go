//2.下面代码能编译通过吗？可以的话，输出什么？
package main

func main() {

	println(DeferTest1(1))
	println(DeferTest2(1))
}

func DeferTest1(i int) (r int) {
	r = i
	defer func() {
		r += 3
	}()
	return r
}

func DeferTest2(i int) (r int) {
	defer func() {
		r += i
	}()
	return 2
}

//参考答案及解析：43。具体解析请看《5 年 Gopher 都不知道的 defer 细节，你别再掉进坑里！》。https://mp.weixin.qq.com/s/Hm8MdrqYgCQPQ4A1nrv4sw
