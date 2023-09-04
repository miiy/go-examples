package main
//1.下面列举的是 recover() 的几种调用方式，哪些是正确的？
//A.
//func main() {
//	recover()
//	panic(1)
//}
//B.
//func main() {
//	defer recover()
//	panic(1)
//}
//C.
//func main() {
//	defer func() {
//		recover()
//	}()
//	panic(1)
//}
//D.
//func main() {
//	defer func() {
//		defer func() {
//			recover()
//		}()
//	}()
//	panic(1)
//}
//
//参考答案及解析：C。recover() 必须在 defer() 函数中直接调用才有效。上面其他几种情况调用都是无效的：直接调用 recover()、在 defer() 中直接调用 recover() 和 defer() 调用时多层嵌套。
