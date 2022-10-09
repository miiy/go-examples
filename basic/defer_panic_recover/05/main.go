package main

func main() {
	defer println("in main")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()
	panic("panic once")
}

//in main
//panic: panic once
//panic: panic again
//panic: panic again and again
