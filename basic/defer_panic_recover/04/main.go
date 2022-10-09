package main

func main() {
	defer println("in main")
	if err := recover(); err != nil {
		println(err)
	}

	panic("unknown err")
}

// in main
// panic: unknown err

// recover 只有在发生 panic 之后调用才会生效