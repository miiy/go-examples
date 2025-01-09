package main

import "fmt"

type People struct {
}

func (p *People) Name() {
	fmt.Println("zhangsan")
}

func main() {
	nf := (*People).Name
	nf(&People{})

	new(People).Name()

	//&People{}.Name()
	(&People{}).Name()
}

// Output:
// zhangsan
// zhangsan
// zhangsan
