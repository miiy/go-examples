package main
//1.关于 cap 函数适用下面哪些类型？
//A. 数组；
//B. channel;
//C. map；
//D. slice；
//
//参考答案即解析：ABD。cap() 函数的作用：
//
//arry 返回数组的元素个数；
//slice 返回 slice 的最大容量；
//channel 返回 channel 的容量；