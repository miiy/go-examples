package main

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `
		<!DOCTYPE html>
		<html>
		<head>
		  <meta charset="utf-8">
		  <meta name="viewport" content="width=device-width">
		</head>
		<body>
			<div id="example"></div>
			<button onclick="conversation()">btn</button>
		<script>
		
		function conversation()
		{
			source = new EventSource('http://localhost:8081/conversation');
		  var div = document.getElementById('example');
		  
		  source.onopen = function (event) {
			div.innerHTML += '<p>Connection open ...</p>';
		  };
		  
		  source.onerror = function (event) {
			div.innerHTML += '<p>Connection close.</p>';
		  };
		  
		//   source.addEventListener('connecttime', function (event) {
		//     div.innerHTML += ('<p>Start time: ' + event.data + '</p>');
		//   }, false);
		  
		  source.onmessage = function (event) {
			if (event.data == "[DONE]") {
				source.close();
				return
			}
			
			// div.innerHTML += event.data.substring(1, a.length - 1);
			let msg =  event.data
			div.innerHTML += msg.substr(1, msg.length-2);
			console.log('connection is closed');
		  };
		}
		
		</script>
		</body>
		</html>
		`)
	})
	http.HandleFunc("/conversation", forwardRequest)
	http.ListenAndServe(":8081", nil)
}

func forwardRequest(w http.ResponseWriter, r *http.Request) {
	reqBody := bytes.NewReader([]byte(`
	{
		"message": "hello",
		"stream": true
	}
	`))
	client := &http.Client{}
	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "http://localhost:3000/conversation", reqBody)
	if err != nil {
		log.Println(err)
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	// header
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/event-stream")
	// status
	w.WriteHeader(resp.StatusCode)

	// flusher
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		// log.Printf("read: %s\n", line)
		if err != nil {
			if err != io.EOF {
				log.Println("err:", err)
			}

			break
		}
		w.Write(line)
		flusher.Flush()
	}
	return
}
