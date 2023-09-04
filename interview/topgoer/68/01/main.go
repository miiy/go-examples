//1.下面代码有什么问题吗？
package main

func main()  {

	for i:=0;i<10 ;i++  {
	loop:
		println(i)
	}
	goto loop
}

//参考答案及解析：goto 不能跳转到其他函数或者内层代码。编译报错：
//
//goto loop jumps into block starting at
