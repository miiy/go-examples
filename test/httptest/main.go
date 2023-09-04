package main

import (
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	name := r.Form.Get("name")
	io.WriteString(w, "hello, "+name)
}

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
