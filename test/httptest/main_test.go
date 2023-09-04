package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// new request
	postData := url.Values{}
	postData.Add("name", "zhangsan")

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(postData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// new response recoder
	w := httptest.NewRecorder()

	helloHandler(w, req)

	t.Logf("%s", w.Body.Bytes())
}
