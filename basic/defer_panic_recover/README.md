## defer

defer 会在当前函数返回前执行传入的函数，它经常被用于关闭文件描述符、关闭数据库连接以及解锁资源。

defer的顺序是先进后出

函数的参数会被预先计算


## panic

调用 panic 后会立刻停止当前函数的剩余代码，并在当前 Goroutine 中递归执行调用的 defer

panic 只会触发当前 Goroutine 的 defer

recover 可以中止 panic 造成的程序崩溃

recover 只有在 defer 中调用才会生效

## panic and recover




## Articles

[Go Defer Simplified with Practical Visuals](https://blog.learngoprogramming.com/golang-defer-simplified-77d3b2b817ff)

5 More Gotchas of Defer in Go

[5 Gotchas of Defer in Go — Part I](https://blog.learngoprogramming.com/gotchas-of-defer-in-go-1-8d070894cb01)

[5 More Gotchas of Defer in Go — Part II](https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-ii-cc550f6ad9aa)

[5 More Gotchas of Defer in Go — Part III](https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-iii-36a1ab3d6ef1)
