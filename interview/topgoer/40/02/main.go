package main
//2.下面的代码有什么问题？
//func Stop(stop <-chan bool) {
//	close(stop)
//}
//参考答案及解析：有方向的 channel 不可以被关闭。
