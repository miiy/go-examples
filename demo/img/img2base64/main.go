package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

func main() {
	f := flag.String("f", "", "-f path/to/example.jpg")
	flag.Parse()
	b, err := os.ReadFile(*f)
	if err != nil {
		fmt.Println(err)
	}
	s := base64.StdEncoding.EncodeToString(b)
	fmt.Println(s)
}
