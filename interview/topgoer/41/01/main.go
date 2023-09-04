//1.下面代码编译能通过吗？
package main

func main()
{
	fmt.Println("hello world")
}

//参考答案及解析：编译错误。
//
//syntax error: unexpected semicolon or newline before {
//Go 语言中，大括号不能放在单独的一行。
//
//正确的代码如下：
//
//func main(){
//	fmt.Println("hello world")
//}
//引自：http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/
