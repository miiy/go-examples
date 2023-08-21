package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond

func main()  {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	// 传递一个带有超时的 context 来告诉阻塞函数它超时后应该放弃其工作。
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}

//Output:
//context deadline exceeded
