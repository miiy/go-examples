package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:8082", nil))
	}()

	select {}
}

// 打开 http://localhost:8082/debug/pprof/
