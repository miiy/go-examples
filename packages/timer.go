package main

import (
	"fmt"
	"time"
)

func main()  {
	timer := time.NewTimer(2 * time.Second)
	t := <- timer.C
	for true {
		fmt.Println(t)
	}
}