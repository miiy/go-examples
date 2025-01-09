package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting the server ...")
	listener, err := http.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}

}