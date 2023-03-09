package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()

		fmt.Println(line)
		os.Mkdir(line, 0777)
	}
}
