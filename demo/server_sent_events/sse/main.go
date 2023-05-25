package main

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"time"
)

//go:embed index.html
var indexHTML []byte

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(indexHTML)
	})

	http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "SSE not supported", http.StatusInternalServerError)
			return
		}

		fmt.Println("Request received")

		w.Header().Set("Content-Type", "text/event-stream")

		io.WriteString(w, "retry: 1000\n")
		io.WriteString(w, "event: connecttime\n")
		io.WriteString(w, fmt.Sprintf("data: %s\n\n", time.Now()))

		for i := 0; i < 10; i++ {
			io.WriteString(w, fmt.Sprintf("data: %s\n\n", time.Now()))
			flusher.Flush()
			time.Sleep(1 * time.Second)
		}
	})

	http.ListenAndServe(":8081", nil)
}
