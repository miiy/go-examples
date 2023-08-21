package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	f := flag.String("f", "", "-f /path/to/src.txt")
	flag.Parse()
	fb, err := os.ReadFile(*f)
	b, err := base64.StdEncoding.DecodeString(string(fb))
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("tmp", b, 0644)
	if err != nil {
		log.Println(err)
	}
}
